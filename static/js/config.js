const config = {
    SSL: false,

    get httpURL() {
        return window.location.origin;
    },

    get wsURL() {
        return (window.location.protocol === "https:")
            ? `wss://${window.location.host}`
            : `ws://${window.location.host}`;
    },

    get lobbyWSEndpoint() {
        return `${this.wsURL}/ws/lobby`;
    },

    get liveWSEndpoint() {
        return `${this.wsURL}/ws/live`;
    }
}

Object.freeze(config);
