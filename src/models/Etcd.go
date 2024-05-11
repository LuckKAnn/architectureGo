package models

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var EtcdClient *clientv3.Client

func init() {
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2479"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
}

func InsertData(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = EtcdClient.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return err
	}
	return nil
}

func QueryValue(key string) []byte {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := EtcdClient.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return nil
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		return ev.Value
	}
	return nil
}
