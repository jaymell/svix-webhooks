pub mod cluster;

pub use cluster::RedisClusterConnectionManager;

// use axum::async_trait;
// use redis::ErrorKind;
// use redis::{IntoConnectionInfo, RedisError, aio::ConnectionLike};
// use redis_cluster_async::Client;
// pub use cluster::RedisClusterConnectionManager;

// pub struct SvixManagedConnection {
//     inner: bb8::ManageConnection<Connection = inner::Connection, Error=inner::Error>
// }

// #[async_trait]
// impl bb8::ManageConnection for SvixManagedConnection {
//    type Connection = inner::Connection;
//    type Error = inner::Error;

//     async fn connect(&self) -> Result<Self::Connection, Self::Error> {
//             self.inner.client.get_connection().await
//     }

//     async fn is_valid(
//         &self,
//         conn: &mut bb8::PooledConnection<'_, Self>,
//     ) -> Result<(), Self::Error> {
//         Ok(())
//         // let pong: String = redis::cmd("PING").query_async(conn).await?;
//         // match pong.as_str() {
//         //     "PONG" => Ok(()),
//         //     _ => Err((ErrorKind::ResponseError, "ping request").into()),
//         // }
//     }


//     fn has_broken(&self, _: &mut Self::Connection) -> bool {
//         false
//     }
// } 

