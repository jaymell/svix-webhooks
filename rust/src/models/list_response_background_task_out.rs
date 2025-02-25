// this file is @generated
use serde::{Deserialize, Serialize};

use super::background_task_out::BackgroundTaskOut;

#[derive(Clone, Debug, Default, PartialEq, Deserialize, Serialize)]
pub struct ListResponseBackgroundTaskOut {
    pub data: Vec<BackgroundTaskOut>,

    pub done: bool,

    pub iterator: String,

    #[serde(rename = "prevIterator")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub prev_iterator: Option<String>,
}

impl ListResponseBackgroundTaskOut {
    pub fn new(data: Vec<BackgroundTaskOut>, done: bool, iterator: String) -> Self {
        Self {
            data,
            done,
            iterator,
            prev_iterator: None,
        }
    }
}
