import {
    Configuration,
    BackgroundTaskStatus,
    BackgroundTaskType,
    BackgroundTaskOut,
    ListResponseBackgroundTaskOut,
    BackgroundTasksApi,
} from "../openapi";
import { ListOptions } from "../util";

export interface BackgroundTaskListOptions extends ListOptions {
    status?: BackgroundTaskStatus;
    task?: BackgroundTaskType;
}

export class BackgroundTask {
    private readonly api: BackgroundTasksApi;

    public constructor(config: Configuration) {
        this.api = new BackgroundTasksApi(config);
    }

    public listByEndpoint(
        options?: BackgroundTaskListOptions
    ): Promise<ListResponseBackgroundTaskOut> {
        const iterator = options?.iterator ?? undefined;
        return this.api.listBackgroundTasks({ ...options, iterator });
    }

    public get(taskId: string): Promise<BackgroundTaskOut> {
        return this.api.getBackgroundTask({
            taskId,
        });
    }
}
