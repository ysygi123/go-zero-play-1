package syredis

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"log"
)

var GsyRedis *SyRedis

type SyRedis struct {
	dispatcher *hash.ConsistentHash
}

func InitSyRedis(c cache.CacheConf) {
	GsyRedis = NewSyRedis(c)
}

// NewSyRedis returns a Store.
func NewSyRedis(c cache.CacheConf) *SyRedis {
	if len(c) == 0 || cache.TotalWeights(c) <= 0 {
		log.Fatal("no cache nodes")
	}

	// even if only one node, we chose to use consistent hash,
	// because Store and redis.Redis has different methods.
	dispatcher := hash.NewConsistentHash()
	for _, node := range c {
		cn := node.NewRedis()
		dispatcher.AddWithWeight(cn, node.Weight)
	}

	return &SyRedis{
		dispatcher: dispatcher,
	}
}

func (cs *SyRedis) Decr(key string) (int64, error) {
	return cs.DecrCtx(context.Background(), key)
}

func (cs *SyRedis) DecrCtx(ctx context.Context, key string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.DecrCtx(ctx, key)
}

func (cs *SyRedis) Decrby(key string, decrement int64) (int64, error) {
	return cs.DecrbyCtx(context.Background(), key, decrement)
}

func (cs *SyRedis) DecrbyCtx(ctx context.Context, key string, decrement int64) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.DecrbyCtx(ctx, key, decrement)
}

func (cs *SyRedis) Del(keys ...string) (int, error) {
	return cs.DelCtx(context.Background(), keys...)
}

func (cs *SyRedis) DelCtx(ctx context.Context, keys ...string) (int, error) {
	var val int
	var be errorx.BatchError

	for _, key := range keys {
		node, e := cs.getRedis(key)
		if e != nil {
			be.Add(e)
			continue
		}

		if v, e := node.DelCtx(ctx, key); e != nil {
			be.Add(e)
		} else {
			val += v
		}
	}

	return val, be.Err()
}

func (cs *SyRedis) Eval(script, key string, args ...interface{}) (interface{}, error) {
	return cs.EvalCtx(context.Background(), script, key, args...)
}

func (cs *SyRedis) EvalCtx(ctx context.Context, script, key string, args ...interface{}) (interface{}, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.EvalCtx(ctx, script, []string{key}, args...)
}

func (cs *SyRedis) Exists(key string) (bool, error) {
	return cs.ExistsCtx(context.Background(), key)
}

func (cs *SyRedis) ExistsCtx(ctx context.Context, key string) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.ExistsCtx(ctx, key)
}

func (cs *SyRedis) Expire(key string, seconds int) error {
	return cs.ExpireCtx(context.Background(), key, seconds)
}

func (cs *SyRedis) ExpireCtx(ctx context.Context, key string, seconds int) error {
	node, err := cs.getRedis(key)
	if err != nil {
		return err
	}

	return node.ExpireCtx(ctx, key, seconds)
}

func (cs *SyRedis) Expireat(key string, expireTime int64) error {
	return cs.ExpireatCtx(context.Background(), key, expireTime)
}

func (cs *SyRedis) ExpireatCtx(ctx context.Context, key string, expireTime int64) error {
	node, err := cs.getRedis(key)
	if err != nil {
		return err
	}

	return node.ExpireatCtx(ctx, key, expireTime)
}

func (cs *SyRedis) Get(key string) (string, error) {
	return cs.GetCtx(context.Background(), key)
}

func (cs *SyRedis) GetCtx(ctx context.Context, key string) (string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return "", err
	}

	return node.GetCtx(ctx, key)
}

func (cs *SyRedis) Hdel(key, field string) (bool, error) {
	return cs.HdelCtx(context.Background(), key, field)
}

func (cs *SyRedis) HdelCtx(ctx context.Context, key, field string) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.HdelCtx(ctx, key, field)
}

func (cs *SyRedis) Hexists(key, field string) (bool, error) {
	return cs.HexistsCtx(context.Background(), key, field)
}

func (cs *SyRedis) HexistsCtx(ctx context.Context, key, field string) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.HexistsCtx(ctx, key, field)
}

func (cs *SyRedis) Hget(key, field string) (string, error) {
	return cs.HgetCtx(context.Background(), key, field)
}

func (cs *SyRedis) HgetCtx(ctx context.Context, key, field string) (string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return "", err
	}

	return node.HgetCtx(ctx, key, field)
}

func (cs *SyRedis) Hgetall(key string) (map[string]string, error) {
	return cs.HgetallCtx(context.Background(), key)
}

func (cs *SyRedis) HgetallCtx(ctx context.Context, key string) (map[string]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.HgetallCtx(ctx, key)
}

func (cs *SyRedis) Hincrby(key, field string, increment int) (int, error) {
	return cs.HincrbyCtx(context.Background(), key, field, increment)
}

func (cs *SyRedis) HincrbyCtx(ctx context.Context, key, field string, increment int) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.HincrbyCtx(ctx, key, field, increment)
}

func (cs *SyRedis) Hkeys(key string) ([]string, error) {
	return cs.HkeysCtx(context.Background(), key)
}

func (cs *SyRedis) HkeysCtx(ctx context.Context, key string) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.HkeysCtx(ctx, key)
}

func (cs *SyRedis) Hlen(key string) (int, error) {
	return cs.HlenCtx(context.Background(), key)
}

func (cs *SyRedis) HlenCtx(ctx context.Context, key string) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.HlenCtx(ctx, key)
}

func (cs *SyRedis) Hmget(key string, fields ...string) ([]string, error) {
	return cs.HmgetCtx(context.Background(), key, fields...)
}

func (cs *SyRedis) HmgetCtx(ctx context.Context, key string, fields ...string) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.HmgetCtx(ctx, key, fields...)
}

func (cs *SyRedis) Hset(key, field, value string) error {
	return cs.HsetCtx(context.Background(), key, field, value)
}

func (cs *SyRedis) HsetCtx(ctx context.Context, key, field, value string) error {
	node, err := cs.getRedis(key)
	if err != nil {
		return err
	}

	return node.HsetCtx(ctx, key, field, value)
}

func (cs *SyRedis) Hsetnx(key, field, value string) (bool, error) {
	return cs.HsetnxCtx(context.Background(), key, field, value)
}

func (cs *SyRedis) HsetnxCtx(ctx context.Context, key, field, value string) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.HsetnxCtx(ctx, key, field, value)
}

func (cs *SyRedis) Hmset(key string, fieldsAndValues map[string]string) error {
	return cs.HmsetCtx(context.Background(), key, fieldsAndValues)
}

func (cs *SyRedis) HmsetCtx(ctx context.Context, key string, fieldsAndValues map[string]string) error {
	node, err := cs.getRedis(key)
	if err != nil {
		return err
	}

	return node.HmsetCtx(ctx, key, fieldsAndValues)
}

func (cs *SyRedis) Hvals(key string) ([]string, error) {
	return cs.HvalsCtx(context.Background(), key)
}

func (cs *SyRedis) HvalsCtx(ctx context.Context, key string) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.HvalsCtx(ctx, key)
}

func (cs *SyRedis) Incr(key string) (int64, error) {
	return cs.IncrCtx(context.Background(), key)
}

func (cs *SyRedis) IncrCtx(ctx context.Context, key string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.IncrCtx(ctx, key)
}

func (cs *SyRedis) Incrby(key string, increment int64) (int64, error) {
	return cs.IncrbyCtx(context.Background(), key, increment)
}

func (cs *SyRedis) IncrbyCtx(ctx context.Context, key string, increment int64) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.IncrbyCtx(ctx, key, increment)
}

func (cs *SyRedis) Llen(key string) (int, error) {
	return cs.LlenCtx(context.Background(), key)
}

func (cs *SyRedis) LlenCtx(ctx context.Context, key string) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.LlenCtx(ctx, key)
}

func (cs *SyRedis) Lindex(key string, index int64) (string, error) {
	return cs.LindexCtx(context.Background(), key, index)
}

func (cs *SyRedis) LindexCtx(ctx context.Context, key string, index int64) (string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return "", err
	}

	return node.LindexCtx(ctx, key, index)
}

func (cs *SyRedis) Lpop(key string) (string, error) {
	return cs.LpopCtx(context.Background(), key)
}

func (cs *SyRedis) LpopCtx(ctx context.Context, key string) (string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return "", err
	}

	return node.LpopCtx(ctx, key)
}

func (cs *SyRedis) Lpush(key string, values ...interface{}) (int, error) {
	return cs.LpushCtx(context.Background(), key, values...)
}

func (cs *SyRedis) LpushCtx(ctx context.Context, key string, values ...interface{}) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.LpushCtx(ctx, key, values...)
}

func (cs *SyRedis) Lrange(key string, start, stop int) ([]string, error) {
	return cs.LrangeCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) LrangeCtx(ctx context.Context, key string, start, stop int) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.LrangeCtx(ctx, key, start, stop)
}

func (cs *SyRedis) Lrem(key string, count int, value string) (int, error) {
	return cs.LremCtx(context.Background(), key, count, value)
}

func (cs *SyRedis) LremCtx(ctx context.Context, key string, count int, value string) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.LremCtx(ctx, key, count, value)
}

func (cs *SyRedis) Persist(key string) (bool, error) {
	return cs.PersistCtx(context.Background(), key)
}

func (cs *SyRedis) PersistCtx(ctx context.Context, key string) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.PersistCtx(ctx, key)
}

func (cs *SyRedis) Pfadd(key string, values ...interface{}) (bool, error) {
	return cs.PfaddCtx(context.Background(), key, values...)
}

func (cs *SyRedis) PfaddCtx(ctx context.Context, key string, values ...interface{}) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.PfaddCtx(ctx, key, values...)
}

func (cs *SyRedis) Pfcount(key string) (int64, error) {
	return cs.PfcountCtx(context.Background(), key)
}

func (cs *SyRedis) PfcountCtx(ctx context.Context, key string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.PfcountCtx(ctx, key)
}

func (cs *SyRedis) Rpush(key string, values ...interface{}) (int, error) {
	return cs.RpushCtx(context.Background(), key, values...)
}

func (cs *SyRedis) RpushCtx(ctx context.Context, key string, values ...interface{}) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.RpushCtx(ctx, key, values...)
}

func (cs *SyRedis) Sadd(key string, values ...interface{}) (int, error) {
	return cs.SaddCtx(context.Background(), key, values...)
}

func (cs *SyRedis) SaddCtx(ctx context.Context, key string, values ...interface{}) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.SaddCtx(ctx, key, values...)
}

func (cs *SyRedis) Scard(key string) (int64, error) {
	return cs.ScardCtx(context.Background(), key)
}

func (cs *SyRedis) ScardCtx(ctx context.Context, key string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ScardCtx(ctx, key)
}

func (cs *SyRedis) Set(key, value string) error {
	return cs.SetCtx(context.Background(), key, value)
}

func (cs *SyRedis) SetCtx(ctx context.Context, key, value string) error {
	node, err := cs.getRedis(key)
	if err != nil {
		return err
	}

	return node.SetCtx(ctx, key, value)
}

func (cs *SyRedis) Setex(key, value string, seconds int) error {
	return cs.SetexCtx(context.Background(), key, value, seconds)
}

func (cs *SyRedis) SetexCtx(ctx context.Context, key, value string, seconds int) error {
	node, err := cs.getRedis(key)
	if err != nil {
		return err
	}

	return node.SetexCtx(ctx, key, value, seconds)
}

func (cs *SyRedis) Setnx(key, value string) (bool, error) {
	return cs.SetnxCtx(context.Background(), key, value)
}

func (cs *SyRedis) SetnxCtx(ctx context.Context, key, value string) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.SetnxCtx(ctx, key, value)
}

func (cs *SyRedis) SetnxEx(key, value string, seconds int) (bool, error) {
	return cs.SetnxExCtx(context.Background(), key, value, seconds)
}

func (cs *SyRedis) SetnxExCtx(ctx context.Context, key, value string, seconds int) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.SetnxExCtx(ctx, key, value, seconds)
}

func (cs *SyRedis) GetSet(key, value string) (string, error) {
	return cs.GetSetCtx(context.Background(), key, value)
}

func (cs *SyRedis) GetSetCtx(ctx context.Context, key, value string) (string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return "", err
	}

	return node.GetSetCtx(ctx, key, value)
}

func (cs *SyRedis) Sismember(key string, value interface{}) (bool, error) {
	return cs.SismemberCtx(context.Background(), key, value)
}

func (cs *SyRedis) SismemberCtx(ctx context.Context, key string, value interface{}) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.SismemberCtx(ctx, key, value)
}

func (cs *SyRedis) Smembers(key string) ([]string, error) {
	return cs.SmembersCtx(context.Background(), key)
}

func (cs *SyRedis) SmembersCtx(ctx context.Context, key string) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.SmembersCtx(ctx, key)
}

func (cs *SyRedis) Spop(key string) (string, error) {
	return cs.SpopCtx(context.Background(), key)
}

func (cs *SyRedis) SpopCtx(ctx context.Context, key string) (string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return "", err
	}

	return node.SpopCtx(ctx, key)
}

func (cs *SyRedis) Srandmember(key string, count int) ([]string, error) {
	return cs.SrandmemberCtx(context.Background(), key, count)
}

func (cs *SyRedis) SrandmemberCtx(ctx context.Context, key string, count int) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.SrandmemberCtx(ctx, key, count)
}

func (cs *SyRedis) Srem(key string, values ...interface{}) (int, error) {
	return cs.SremCtx(context.Background(), key, values...)
}

func (cs *SyRedis) SremCtx(ctx context.Context, key string, values ...interface{}) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.SremCtx(ctx, key, values...)
}

func (cs *SyRedis) Sscan(key string, cursor uint64, match string, count int64) (
	keys []string, cur uint64, err error) {
	return cs.SscanCtx(context.Background(), key, cursor, match, count)
}

func (cs *SyRedis) SscanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (
	keys []string, cur uint64, err error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, 0, err
	}

	return node.SscanCtx(ctx, key, cursor, match, count)
}

func (cs *SyRedis) Ttl(key string) (int, error) {
	return cs.TtlCtx(context.Background(), key)
}

func (cs *SyRedis) TtlCtx(ctx context.Context, key string) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.TtlCtx(ctx, key)
}

func (cs *SyRedis) Zadd(key string, score int64, value string) (bool, error) {
	return cs.ZaddCtx(context.Background(), key, score, value)
}

func (cs *SyRedis) ZaddCtx(ctx context.Context, key string, score int64, value string) (bool, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return false, err
	}

	return node.ZaddCtx(ctx, key, score, value)
}

func (cs *SyRedis) Zadds(key string, ps ...redis.Pair) (int64, error) {
	return cs.ZaddsCtx(context.Background(), key, ps...)
}

func (cs *SyRedis) ZaddsCtx(ctx context.Context, key string, ps ...redis.Pair) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZaddsCtx(ctx, key, ps...)
}

func (cs *SyRedis) Zcard(key string) (int, error) {
	return cs.ZcardCtx(context.Background(), key)
}

func (cs *SyRedis) ZcardCtx(ctx context.Context, key string) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZcardCtx(ctx, key)
}

func (cs *SyRedis) Zcount(key string, start, stop int64) (int, error) {
	return cs.ZcountCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZcountCtx(ctx context.Context, key string, start, stop int64) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZcountCtx(ctx, key, start, stop)
}

func (cs *SyRedis) Zincrby(key string, increment int64, field string) (int64, error) {
	return cs.ZincrbyCtx(context.Background(), key, increment, field)
}

func (cs *SyRedis) ZincrbyCtx(ctx context.Context, key string, increment int64, field string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZincrbyCtx(ctx, key, increment, field)
}

func (cs *SyRedis) Zrank(key, field string) (int64, error) {
	return cs.ZrankCtx(context.Background(), key, field)
}

func (cs *SyRedis) ZrankCtx(ctx context.Context, key, field string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZrankCtx(ctx, key, field)
}

func (cs *SyRedis) Zrange(key string, start, stop int64) ([]string, error) {
	return cs.ZrangeCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZrangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.ZrangeCtx(ctx, key, start, stop)
}

func (cs *SyRedis) ZrangeWithScores(key string, start, stop int64) ([]redis.Pair, error) {
	return cs.ZrangeWithScoresCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZrangeWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.ZrangeWithScoresCtx(ctx, key, start, stop)
}

func (cs *SyRedis) ZrangebyscoreWithScores(key string, start, stop int64) ([]redis.Pair, error) {
	return cs.ZrangebyscoreWithScoresCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.ZrangebyscoreWithScoresCtx(ctx, key, start, stop)
}

func (cs *SyRedis) ZrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	return cs.ZrangebyscoreWithScoresAndLimitCtx(context.Background(), key, start, stop, page, size)
}

func (cs *SyRedis) ZrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.ZrangebyscoreWithScoresAndLimitCtx(ctx, key, start, stop, page, size)
}

func (cs *SyRedis) Zrem(key string, values ...interface{}) (int, error) {
	return cs.ZremCtx(context.Background(), key, values...)
}

func (cs *SyRedis) ZremCtx(ctx context.Context, key string, values ...interface{}) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZremCtx(ctx, key, values...)
}

func (cs *SyRedis) Zremrangebyrank(key string, start, stop int64) (int, error) {
	return cs.ZremrangebyrankCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZremrangebyrankCtx(ctx context.Context, key string, start, stop int64) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZremrangebyrankCtx(ctx, key, start, stop)
}

func (cs *SyRedis) Zremrangebyscore(key string, start, stop int64) (int, error) {
	return cs.ZremrangebyscoreCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZremrangebyscoreCtx(ctx context.Context, key string, start, stop int64) (int, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZremrangebyscoreCtx(ctx, key, start, stop)
}

func (cs *SyRedis) Zrevrange(key string, start, stop int64) ([]string, error) {
	return cs.ZrevrangeCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZrevrangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.ZrevrangeCtx(ctx, key, start, stop)
}

func (cs *SyRedis) ZrevrangebyscoreWithScores(key string, start, stop int64) ([]redis.Pair, error) {
	return cs.ZrevrangebyscoreWithScoresCtx(context.Background(), key, start, stop)
}

func (cs *SyRedis) ZrevrangebyscoreWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]redis.Pair, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.ZrevrangebyscoreWithScoresCtx(ctx, key, start, stop)
}

func (cs *SyRedis) ZrevrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	return cs.ZrevrangebyscoreWithScoresAndLimitCtx(context.Background(), key, start, stop, page, size)
}

func (cs *SyRedis) ZrevrangebyscoreWithScoresAndLimitCtx(ctx context.Context, key string, start, stop int64, page, size int) (
	[]redis.Pair, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return nil, err
	}

	return node.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, key, start, stop, page, size)
}

func (cs *SyRedis) Zrevrank(key, field string) (int64, error) {
	return cs.ZrevrankCtx(context.Background(), key, field)
}

func (cs *SyRedis) ZrevrankCtx(ctx context.Context, key, field string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZrevrankCtx(ctx, key, field)
}

func (cs *SyRedis) Zscore(key, value string) (int64, error) {
	return cs.ZscoreCtx(context.Background(), key, value)
}

func (cs *SyRedis) ZscoreCtx(ctx context.Context, key, value string) (int64, error) {
	node, err := cs.getRedis(key)
	if err != nil {
		return 0, err
	}

	return node.ZscoreCtx(ctx, key, value)
}

func (cs *SyRedis) getRedis(key string) (*redis.Redis, error) {
	val, ok := cs.dispatcher.Get(key)
	if !ok {
		return nil, kv.ErrNoRedisNode
	}

	return val.(*redis.Redis), nil
}

func (cs *SyRedis) GetRedis(key string) (*redis.Redis, error) {
	return cs.getRedis(key)
}
