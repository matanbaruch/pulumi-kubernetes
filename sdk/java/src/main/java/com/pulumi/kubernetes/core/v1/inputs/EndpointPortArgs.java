// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * EndpointPort is a tuple that describes a single port. Deprecated: This API is deprecated in v1.33+.
 * 
 */
public final class EndpointPortArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointPortArgs Empty = new EndpointPortArgs();

    /**
     * The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:
     * 
     * * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).
     * 
     * * Kubernetes-defined prefixed names:
     *   * &#39;kubernetes.io/h2c&#39; - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-
     *   * &#39;kubernetes.io/ws&#39;  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
     *   * &#39;kubernetes.io/wss&#39; - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455
     * 
     * * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
     * 
     */
    @Import(name="appProtocol")
    private @Nullable Output<String> appProtocol;

    /**
     * @return The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:
     * 
     * * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).
     * 
     * * Kubernetes-defined prefixed names:
     *   * &#39;kubernetes.io/h2c&#39; - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-
     *   * &#39;kubernetes.io/ws&#39;  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
     *   * &#39;kubernetes.io/wss&#39; - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455
     * 
     * * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
     * 
     */
    public Optional<Output<String>> appProtocol() {
        return Optional.ofNullable(this.appProtocol);
    }

    /**
     * The name of this port.  This must match the &#39;name&#39; field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of this port.  This must match the &#39;name&#39; field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The port number of the endpoint.
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return The port number of the endpoint.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    /**
     * The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    private EndpointPortArgs() {}

    private EndpointPortArgs(EndpointPortArgs $) {
        this.appProtocol = $.appProtocol;
        this.name = $.name;
        this.port = $.port;
        this.protocol = $.protocol;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointPortArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointPortArgs $;

        public Builder() {
            $ = new EndpointPortArgs();
        }

        public Builder(EndpointPortArgs defaults) {
            $ = new EndpointPortArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param appProtocol The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:
         * 
         * * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).
         * 
         * * Kubernetes-defined prefixed names:
         *   * &#39;kubernetes.io/h2c&#39; - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-
         *   * &#39;kubernetes.io/ws&#39;  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
         *   * &#39;kubernetes.io/wss&#39; - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455
         * 
         * * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
         * 
         * @return builder
         * 
         */
        public Builder appProtocol(@Nullable Output<String> appProtocol) {
            $.appProtocol = appProtocol;
            return this;
        }

        /**
         * @param appProtocol The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:
         * 
         * * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).
         * 
         * * Kubernetes-defined prefixed names:
         *   * &#39;kubernetes.io/h2c&#39; - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-
         *   * &#39;kubernetes.io/ws&#39;  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
         *   * &#39;kubernetes.io/wss&#39; - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455
         * 
         * * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
         * 
         * @return builder
         * 
         */
        public Builder appProtocol(String appProtocol) {
            return appProtocol(Output.of(appProtocol));
        }

        /**
         * @param name The name of this port.  This must match the &#39;name&#39; field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of this port.  This must match the &#39;name&#39; field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param port The port number of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port number of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        public EndpointPortArgs build() {
            if ($.port == null) {
                throw new MissingRequiredPropertyException("EndpointPortArgs", "port");
            }
            return $;
        }
    }

}
