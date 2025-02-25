// this file is @generated
package com.svix.kotlin

import com.svix.kotlin.exceptions.ApiException
import com.svix.kotlin.internal.apis.StatisticsApi
import com.svix.kotlin.models.AggregateEventTypesOut
import com.svix.kotlin.models.AppUsageStatsIn
import com.svix.kotlin.models.AppUsageStatsOut

class Statistics internal constructor(token: String, options: SvixOptions) {
    private val api = StatisticsApi(options.serverUrl)

    init {
        api.accessToken = token
        api.userAgent = options.getUA()
        options.initialRetryDelayMillis?.let { api.initialRetryDelayMillis = it }
        options.numRetries?.let { api.numRetries = it }
    }

    /**
     * Creates a background task to calculate the message destinations for all applications in the
     * environment.
     *
     * Note that this endpoint is asynchronous. You will need to poll the `Get Background Task`
     * endpoint to retrieve the results of the operation.
     */
    suspend fun aggregateAppStats(
        appUsageStatsIn: AppUsageStatsIn,
        options: PostOptions = PostOptions(),
    ): AppUsageStatsOut {
        try {
            return api.v1StatisticsAggregateAppStats(appUsageStatsIn, options.idempotencyKey)
        } catch (e: Exception) {
            throw ApiException.wrap(e)
        }
    }

    /**
     * Creates a background task to calculate the listed event types for all apps in the
     * organization.
     *
     * Note that this endpoint is asynchronous. You will need to poll the `Get Background Task`
     * endpoint to retrieve the results of the operation.
     */
    suspend fun aggregateEventTypes(): AggregateEventTypesOut {
        try {
            return api.v1StatisticsAggregateEventTypes()
        } catch (e: Exception) {
            throw ApiException.wrap(e)
        }
    }
}
