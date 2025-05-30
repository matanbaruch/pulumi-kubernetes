// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// ResourceStatus represents the status of a single resource allocated to a Pod.
    /// </summary>
    [OutputType]
    public sealed class ResourceStatusPatch
    {
        /// <summary>
        /// Name of the resource. Must be unique within the pod and in case of non-DRA resource, match one of the resources from the pod spec. For DRA resources, the value must be "claim:&lt;claim_name&gt;/&lt;request&gt;". When this status is reported about a container, the "claim_name" and "request" must match one of the claims of this container.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of unique resources health. Each element in the list contains an unique resource ID and its health. At a minimum, for the lifetime of a Pod, resource ID must uniquely identify the resource allocated to the Pod on the Node. If other Pod on the same Node reports the status with the same resource ID, it must be the same resource they share. See ResourceID type definition for a specific format it has in various use cases.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceHealthPatch> Resources;

        [OutputConstructor]
        private ResourceStatusPatch(
            string name,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceHealthPatch> resources)
        {
            Name = name;
            Resources = resources;
        }
    }
}
