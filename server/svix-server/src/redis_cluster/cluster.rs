use axum::async_trait;
use redis::ErrorKind;
use redis::{IntoConnectionInfo, RedisError, aio::ConnectionLike};
use redis_cluster_async::Client;


#[derive(Clone)]
pub struct RedisClusterConnectionManager {
    client: Client,
}

impl RedisClusterConnectionManager {
    pub fn new<T: IntoConnectionInfo>(info: Vec<T>) -> Result<RedisClusterConnectionManager, RedisError> {
        Ok(RedisClusterConnectionManager {
            client: Client::open(info)?,
        })
    }
}

#[async_trait]
impl bb8::ManageConnection for RedisClusterConnectionManager {
    type Connection = redis_cluster_async::Connection;
    type Error = RedisError;

    async fn connect(&self) -> Result<Self::Connection, Self::Error> {
        self.client.get_connection().await
    }

    async fn is_valid(
        &self,
        conn: &mut bb8::PooledConnection<'_, Self>,
    ) -> Result<(), Self::Error> {
        Ok(())
        // let pong: String = redis::cmd("PING").query_async(conn).await?;
        // match pong.as_str() {
        //     "PONG" => Ok(()),
        //     _ => Err((ErrorKind::ResponseError, "ping request").into()),
        // }
    }

    fn has_broken(&self, _: &mut Self::Connection) -> bool {
        false
    }
}