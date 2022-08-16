package server

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/the-zion/matrix-core/app/message/service/internal/conf"
	"github.com/the-zion/matrix-core/app/message/service/internal/service"
)

type ArticleReviewMqConsumerServer struct {
	c  rocketmq.PushConsumer
	ms *service.MessageService
}

func NewArticleReviewMqConsumerServer(conf *conf.Server, messageService *service.MessageService, logger log.Logger) *ArticleReviewMqConsumerServer {
	l := log.NewHelper(log.With(logger, "server", "message/server/rocketmq-article-review-consumer"))
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.CreationMq.ArticleReview.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.CreationMq.ServerAddress})),
		consumer.WithCredentials(primitive.Credentials{
			SecretKey: conf.CreationMq.SecretKey,
			AccessKey: conf.CreationMq.AccessKey,
		}),
		consumer.WithInstance("creation"),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithNamespace(conf.CreationMq.NameSpace),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		l.Fatalf("init consumer error: %v", err)
	}

	delayLevel := 3
	err = c.Subscribe("article_review", consumer.MessageSelector{}, MqRecovery(func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		concurrentCtx, _ := primitive.GetConcurrentlyCtx(ctx)
		concurrentCtx.DelayLevelWhenNextConsume = delayLevel

		var err error
		m := map[string]interface{}{}
		err = json.Unmarshal(msgs[0].Body, &m)
		if err != nil {
			l.Errorf("fail to unmarshal msg: %s", err.Error())
			return consumer.ConsumeRetryLater, nil
		}

		mode := m["Mode"].(string)
		switch mode {
		case "create":
			err = messageService.ToReviewCreateArticle(int32(m["Id"].(float64)), m["Uuid"].(string))
		case "edit":
			err = messageService.ToReviewEditArticle(int32(m["Id"].(float64)), m["Uuid"].(string))
		}

		if err != nil {
			l.Error(err.Error())
			return consumer.ConsumeRetryLater, nil
		}
		return consumer.ConsumeSuccess, nil
	}))
	if err != nil {
		l.Fatalf("consumer subscribe error: %v", err)
	}

	return &ArticleReviewMqConsumerServer{
		c: c,
	}
}

func (s *ArticleReviewMqConsumerServer) Start(_ context.Context) error {
	log.Info("mq article review consumer starting")
	return s.c.Start()
}

func (s *ArticleReviewMqConsumerServer) Stop(_ context.Context) error {
	log.Info("mq article review consumer closing")
	return s.c.Shutdown()
}

type ArticleMqConsumerServer struct {
	c  rocketmq.PushConsumer
	ms *service.MessageService
}

func NewArticleMqConsumerServer(conf *conf.Server, messageService *service.MessageService, logger log.Logger) *ArticleMqConsumerServer {
	l := log.NewHelper(log.With(logger, "server", "message/server/rocketmq-article-consumer"))
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.CreationMq.Article.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.CreationMq.ServerAddress})),
		consumer.WithCredentials(primitive.Credentials{
			SecretKey: conf.CreationMq.SecretKey,
			AccessKey: conf.CreationMq.AccessKey,
		}),
		consumer.WithInstance("creation"),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithNamespace(conf.CreationMq.NameSpace),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		l.Fatalf("init consumer error: %v", err)
	}

	delayLevel := 3
	err = c.Subscribe("article", consumer.MessageSelector{}, MqRecovery(func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		concurrentCtx, _ := primitive.GetConcurrentlyCtx(ctx)
		concurrentCtx.DelayLevelWhenNextConsume = delayLevel

		var err error
		m := map[string]interface{}{}
		err = json.Unmarshal(msgs[0].Body, &m)
		if err != nil {
			l.Errorf("fail to unmarshal msg: %s", err.Error())
			return consumer.ConsumeRetryLater, nil
		}

		mode := m["mode"].(string)
		switch mode {
		case "create_article_db_cache_and_search":
			err = messageService.CreateArticleDbCacheAndSearch(ctx, int32(m["id"].(float64)), int32(m["auth"].(float64)), m["uuid"].(string))
		case "edit_article_cos_and_search":
			err = messageService.EditArticleCosAndSearch(ctx, int32(m["id"].(float64)), int32(m["auth"].(float64)), m["uuid"].(string))
		case "delete_article_cache_and_search":
			err = messageService.DeleteArticleCacheAndSearch(ctx, int32(m["id"].(float64)), m["uuid"].(string))
		}

		if err != nil {
			l.Error(err.Error())
			return consumer.ConsumeRetryLater, nil
		}
		return consumer.ConsumeSuccess, nil
	}))
	if err != nil {
		l.Fatalf("consumer subscribe error: %v", err)
	}

	return &ArticleMqConsumerServer{
		c: c,
	}
}

func (s *ArticleMqConsumerServer) Start(_ context.Context) error {
	log.Info("mq article consumer starting")
	return s.c.Start()
}

func (s *ArticleMqConsumerServer) Stop(_ context.Context) error {
	log.Info("mq article consumer closing")
	return s.c.Shutdown()
}

type TalkReviewMqConsumerServer struct {
	c  rocketmq.PushConsumer
	ms *service.MessageService
}

func NewTalkReviewMqConsumerServer(conf *conf.Server, messageService *service.MessageService, logger log.Logger) *TalkReviewMqConsumerServer {
	l := log.NewHelper(log.With(logger, "server", "message/server/rocketmq-talk-review-consumer"))
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.CreationMq.TalkReview.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.CreationMq.ServerAddress})),
		consumer.WithCredentials(primitive.Credentials{
			SecretKey: conf.CreationMq.SecretKey,
			AccessKey: conf.CreationMq.AccessKey,
		}),
		consumer.WithInstance("creation"),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithNamespace(conf.CreationMq.NameSpace),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		l.Fatalf("init consumer error: %v", err)
	}

	delayLevel := 3
	err = c.Subscribe("talk_review", consumer.MessageSelector{}, MqRecovery(func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		concurrentCtx, _ := primitive.GetConcurrentlyCtx(ctx)
		concurrentCtx.DelayLevelWhenNextConsume = delayLevel

		var err error
		m := map[string]interface{}{}
		err = json.Unmarshal(msgs[0].Body, &m)
		if err != nil {
			l.Errorf("fail to unmarshal msg: %s", err.Error())
			return consumer.ConsumeRetryLater, nil
		}

		mode := m["Mode"].(string)
		switch mode {
		case "create":
			err = messageService.ToReviewCreateTalk(int32(m["Id"].(float64)), m["Uuid"].(string))
		case "edit":
			err = messageService.ToReviewEditTalk(int32(m["Id"].(float64)), m["Uuid"].(string))
		}

		if err != nil {
			l.Error(err.Error())
			return consumer.ConsumeRetryLater, nil
		}
		return consumer.ConsumeSuccess, nil
	}))
	if err != nil {
		l.Fatalf("consumer subscribe error: %v", err)
	}

	return &TalkReviewMqConsumerServer{
		c: c,
	}
}

func (s *TalkReviewMqConsumerServer) Start(_ context.Context) error {
	log.Info("mq talk review consumer starting")
	return s.c.Start()
}

func (s *TalkReviewMqConsumerServer) Stop(_ context.Context) error {
	log.Info("mq talk review consumer closing")
	return s.c.Shutdown()
}

type TalkMqConsumerServer struct {
	c  rocketmq.PushConsumer
	ms *service.MessageService
}

func NewTalkMqConsumerServer(conf *conf.Server, messageService *service.MessageService, logger log.Logger) *TalkMqConsumerServer {
	l := log.NewHelper(log.With(logger, "server", "message/server/rocketmq-talk-consumer"))
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.CreationMq.Talk.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.CreationMq.ServerAddress})),
		consumer.WithCredentials(primitive.Credentials{
			SecretKey: conf.CreationMq.SecretKey,
			AccessKey: conf.CreationMq.AccessKey,
		}),
		consumer.WithInstance("creation"),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithNamespace(conf.CreationMq.NameSpace),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		l.Fatalf("init consumer error: %v", err)
	}

	delayLevel := 3
	err = c.Subscribe("talk", consumer.MessageSelector{}, MqRecovery(func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		concurrentCtx, _ := primitive.GetConcurrentlyCtx(ctx)
		concurrentCtx.DelayLevelWhenNextConsume = delayLevel

		var err error
		m := map[string]interface{}{}
		err = json.Unmarshal(msgs[0].Body, &m)
		if err != nil {
			l.Errorf("fail to unmarshal msg: %s", err.Error())
			return consumer.ConsumeRetryLater, nil
		}

		mode := m["mode"].(string)
		switch mode {
		case "create_talk_db_cache_and_search":
			err = messageService.CreateTalkDbCacheAndSearch(ctx, int32(m["id"].(float64)), int32(m["auth"].(float64)), m["uuid"].(string))
		case "edit_talk_cos_and_search":
			err = messageService.EditTalkCosAndSearch(ctx, int32(m["id"].(float64)), int32(m["auth"].(float64)), m["uuid"].(string))
		case "delete_talk_cache_and_search":
			err = messageService.DeleteTalkCacheAndSearch(ctx, int32(m["id"].(float64)), m["uuid"].(string))
		}

		if err != nil {
			l.Error(err.Error())
			return consumer.ConsumeRetryLater, nil
		}
		return consumer.ConsumeSuccess, nil
	}))
	if err != nil {
		l.Fatalf("consumer subscribe error: %v", err)
	}

	return &TalkMqConsumerServer{
		c: c,
	}
}

func (s *TalkMqConsumerServer) Start(_ context.Context) error {
	log.Info("mq talk consumer starting")
	return s.c.Start()
}

func (s *TalkMqConsumerServer) Stop(_ context.Context) error {
	log.Info("mq talk consumer closing")
	return s.c.Shutdown()
}

type ColumnReviewMqConsumerServer struct {
	c  rocketmq.PushConsumer
	ms *service.MessageService
}

func NewColumnReviewMqConsumerServer(conf *conf.Server, messageService *service.MessageService, logger log.Logger) *ColumnReviewMqConsumerServer {
	l := log.NewHelper(log.With(logger, "server", "message/server/rocketmq-column-review-consumer"))
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.CreationMq.ColumnReview.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.CreationMq.ServerAddress})),
		consumer.WithCredentials(primitive.Credentials{
			SecretKey: conf.CreationMq.SecretKey,
			AccessKey: conf.CreationMq.AccessKey,
		}),
		consumer.WithInstance("creation"),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithNamespace(conf.CreationMq.NameSpace),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		l.Fatalf("init consumer error: %v", err)
	}

	delayLevel := 3
	err = c.Subscribe("column_review", consumer.MessageSelector{}, MqRecovery(func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		concurrentCtx, _ := primitive.GetConcurrentlyCtx(ctx)
		concurrentCtx.DelayLevelWhenNextConsume = delayLevel

		var err error
		m := map[string]interface{}{}
		err = json.Unmarshal(msgs[0].Body, &m)
		if err != nil {
			l.Errorf("fail to unmarshal msg: %s", err.Error())
			return consumer.ConsumeRetryLater, nil
		}

		mode := m["Mode"].(string)
		switch mode {
		case "create":
			err = messageService.ToReviewCreateColumn(int32(m["Id"].(float64)), m["Uuid"].(string))
		case "edit":
			err = messageService.ToReviewEditColumn(int32(m["Id"].(float64)), m["Uuid"].(string))
		}

		if err != nil {
			l.Error(err.Error())
			return consumer.ConsumeRetryLater, nil
		}
		return consumer.ConsumeSuccess, nil
	}))
	if err != nil {
		l.Fatalf("consumer subscribe error: %v", err)
	}

	return &ColumnReviewMqConsumerServer{
		c: c,
	}
}

func (s *ColumnReviewMqConsumerServer) Start(_ context.Context) error {
	log.Info("mq talk review consumer starting")
	return s.c.Start()
}

func (s *ColumnReviewMqConsumerServer) Stop(_ context.Context) error {
	log.Info("mq talk review consumer closing")
	return s.c.Shutdown()
}

type ColumnMqConsumerServer struct {
	c  rocketmq.PushConsumer
	ms *service.MessageService
}

func NewColumnMqConsumerServer(conf *conf.Server, messageService *service.MessageService, logger log.Logger) *ColumnMqConsumerServer {
	l := log.NewHelper(log.With(logger, "server", "message/server/rocketmq-column-consumer"))
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.CreationMq.Column.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.CreationMq.ServerAddress})),
		consumer.WithCredentials(primitive.Credentials{
			SecretKey: conf.CreationMq.SecretKey,
			AccessKey: conf.CreationMq.AccessKey,
		}),
		consumer.WithInstance("creation"),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithNamespace(conf.CreationMq.NameSpace),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		l.Fatalf("init consumer error: %v", err)
	}

	delayLevel := 3
	err = c.Subscribe("column", consumer.MessageSelector{}, MqRecovery(func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		concurrentCtx, _ := primitive.GetConcurrentlyCtx(ctx)
		concurrentCtx.DelayLevelWhenNextConsume = delayLevel

		var err error
		m := map[string]interface{}{}
		err = json.Unmarshal(msgs[0].Body, &m)
		if err != nil {
			l.Errorf("fail to unmarshal msg: %s", err.Error())
			return consumer.ConsumeRetryLater, nil
		}

		mode := m["mode"].(string)
		switch mode {
		case "create_column_db_cache_and_search":
			err = messageService.CreateColumnDbCacheAndSearch(ctx, int32(m["id"].(float64)), int32(m["auth"].(float64)), m["uuid"].(string))
		case "edit_column_cos_and_search":
			err = messageService.EditColumnCosAndSearch(ctx, int32(m["id"].(float64)), int32(m["auth"].(float64)), m["uuid"].(string))
		case "delete_column_cache_and_search":
			err = messageService.DeleteColumnCacheAndSearch(ctx, int32(m["id"].(float64)), m["uuid"].(string))
		}

		if err != nil {
			l.Error(err.Error())
			return consumer.ConsumeRetryLater, nil
		}
		return consumer.ConsumeSuccess, nil
	}))
	if err != nil {
		l.Fatalf("consumer subscribe error: %v", err)
	}

	return &ColumnMqConsumerServer{
		c: c,
	}
}

func (s *ColumnMqConsumerServer) Start(_ context.Context) error {
	log.Info("mq column consumer starting")
	return s.c.Start()
}

func (s *ColumnMqConsumerServer) Stop(_ context.Context) error {
	log.Info("mq column consumer closing")
	return s.c.Shutdown()
}
