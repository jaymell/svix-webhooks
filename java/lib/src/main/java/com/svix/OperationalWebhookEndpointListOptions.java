package com.svix;

import com.svix.models.Ordering;

public class OperationalWebhookEndpointListOptions extends ListOptions {
    private Ordering order;

    public void setOrder(final Ordering order) {
        this.order = order;
    }

    public Ordering getOrder() {
        return this.order;
    }
}
