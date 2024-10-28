package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

//task:
//Напишите кешированную пакетную реализацию для провайдера постов из начала урока.
//Есть провайдер одного поста, который всегда ходит напрямую в базу:
//
//Требования к реализации:
// 1. Если хотя бы один запрос к PostProvider завершается ошибкой, эту ошибку надо вернуть вызывающей стороне.
// 2. Если очередь запросов к PostProvider превышает n, надо сразу же возвращать ошибку errors.New("too many requests").
// 3. Если одновременно вызываются BatchPostProvider с одними и теми же ключами,
//    результаты первого вызова должны переиспользоваться во всех вызывающих горутинах.

type PostID string

type Post struct {
	ID      PostID
	Content string
}

type PostProvider interface {
	GetPost(ctx context.Context, postID PostID) (*Post, error)
}

type MapProvider struct {
	bucket     *semaphore.Weighted
	storageMap map[PostID]Post
	//using sync.Map would be a bad solution
	//because it has a bad optimisation. More over here we use a singleflight.
	sg *singleflight.Group
}

func (m *MapProvider) GetPost(ctx context.Context, postID PostID) (*Post, error) {
	if !m.bucket.TryAcquire(1) {
		return nil, NewErrTooManyRequests()
	}

	v, err, _ := m.sg.Do(string(postID), func() (interface{}, error) {
		p, ok := m.storageMap[postID]
		if !ok {
			return nil, NewErrNotFound()
		} else {
			return &p, nil
		}
	})
	if err != nil {
		return nil, err
	}

	post := v.(*Post)
	return post, nil
}

func NewMapProvider(bucketDepth int) *MapProvider {
	newMap := map[PostID]Post{}
	newMap[PostID("id1")] = Post{
		ID:      PostID("id1"),
		Content: "content 1",
	}
	newMap[PostID("id2")] = Post{
		ID:      PostID("id2"),
		Content: "content 2",
	}
	newMap[PostID("id3")] = Post{
		ID:      PostID("id3"),
		Content: "content 3",
	}
	return &MapProvider{
		bucket:     semaphore.NewWeighted(int64(bucketDepth)),
		storageMap: newMap,
		sg:         &singleflight.Group{},
	}
}

func main() {
	bucketDepth := 10
	provider := NewMapProvider(bucketDepth)

	//no err
	errGrOk, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < bucketDepth/2; i++ {
		gID := i
		errGrOk.Go(func() error {
			p, err := provider.GetPost(ctx, "id1")
			if err != nil {
				return err
			}
			fmt.Printf("NOERR gID = %v, post: %v\n", gID, p)
			return nil
		})
	}
	fmt.Printf("errGrOk (have to be nil): %v\n", errGrOk.Wait())

	//err TooManyRequests:
	errGrTMR, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < bucketDepth*2; i++ {
		gID := i
		errGrTMR.Go(func() error {
			p, err := provider.GetPost(ctx, "id2")
			if err != nil {
				return err
			}
			fmt.Printf("TMR ERR gID = %v, post: %v\n", gID, p)
			return nil
		})
	}
	fmt.Printf("errGrTMR (have to be a \"ToManyRequests\"): %v\n", errGrTMR.Wait())

	//err NotFound
	errGrNF, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < bucketDepth/2; i++ {
		gID := i
		errGrNF.Go(func() error {
			p, err := provider.GetPost(ctx, "someIdWitchNotExists")
			if err != nil {
				return err
			}
			fmt.Printf("NF ERR (this line have not to be printed) gID = %v, post: %v\n", gID, p)
			return nil
		})
	}
	fmt.Printf("errGrNF (have to be a \"NotFound\"): %v\n", errGrNF.Wait())
}
