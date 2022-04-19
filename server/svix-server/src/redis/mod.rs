mod cluster;

use bb8::Pool;
use bb8_redis::RedisConnectionManager;
pub use cluster::RedisClusterConnectionManager;

use axum::async_trait;
use redis::aio::ConnectionLike;

pub struct BoxedConnectionLike(Box<dyn ConnectionLike + Send>);

impl ConnectionLike for BoxedConnectionLike {
    fn req_packed_command<'a>(
        &'a mut self,
        cmd: &'a redis::Cmd,
    ) -> redis::RedisFuture<'a, redis::Value> {
        self.0.req_packed_command(cmd)
    }

    fn req_packed_commands<'a>(
        &'a mut self,
        cmd: &'a redis::Pipeline,
        offset: usize,
        count: usize,
    ) -> redis::RedisFuture<'a, Vec<redis::Value>> {
        self.0.req_packed_commands(cmd, offset, count)
    }

    fn get_db(&self) -> i64 {
        self.0.get_db()
    }
}

#[async_trait]
pub trait PoolLike {
    async fn get(&self) -> BoxedConnectionLike;
}

#[derive(Clone, Debug)]
pub enum SvixRedisPool {
    Clustered(SvixClusteredRedisPool),
    NonClustered(SvixNonClusteredRedisPool),
}

#[async_trait]
impl PoolLike for SvixRedisPool {
    async fn get(&self) -> BoxedConnectionLike {
        match self {
            Self::Clustered(pool) => pool.get().await,
            Self::NonClustered(pool) => pool.get().await,
        }
    }
}

#[derive(Clone, Debug)]
pub struct SvixNonClusteredRedisPool {
    pool: Pool<RedisConnectionManager>,
}

#[async_trait]
impl PoolLike for SvixNonClusteredRedisPool {
    async fn get(&self) -> BoxedConnectionLike {
        let inner_pool = self.pool.get().await.unwrap();
        BoxedConnectionLike(Box::new(*inner_pool))
    }
}

#[derive(Clone, Debug)]
pub struct SvixClusteredRedisPool {
    pool: Pool<RedisClusterConnectionManager>,
}

#[async_trait]
impl PoolLike for SvixClusteredRedisPool {
    async fn get(&self) -> BoxedConnectionLike {
        let inner_pool = self.pool.get().await.unwrap();
        BoxedConnectionLike(Box::new(*inner_pool))
    }
}
