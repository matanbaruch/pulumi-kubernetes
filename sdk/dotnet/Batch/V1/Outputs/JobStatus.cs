// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Batch.V1
{

    /// <summary>
    /// JobStatus represents the current state of a Job.
    /// </summary>
    [OutputType]
    public sealed class JobStatus
    {
        /// <summary>
        /// The number of pending and running pods which are not terminating (without a deletionTimestamp). The value is zero for finished jobs.
        /// </summary>
        public readonly int Active;
        /// <summary>
        /// completedIndexes holds the completed indexes when .spec.completionMode = "Indexed" in a text format. The indexes are represented as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the completed indexes are 1, 3, 4, 5 and 7, they are represented as "1,3-5,7".
        /// </summary>
        public readonly string CompletedIndexes;
        /// <summary>
        /// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. The completion time is set when the job finishes successfully, and only then. The value cannot be updated or removed. The value indicates the same or later point in time as the startTime field.
        /// </summary>
        public readonly string CompletionTime;
        /// <summary>
        /// The latest available observations of an object's current state. When a Job fails, one of the conditions will have type "Failed" and status true. When a Job is suspended, one of the conditions will have type "Suspended" and status true; when the Job is resumed, the status of this condition will become false. When a Job is completed, one of the conditions will have type "Complete" and status true.
        /// 
        /// A job is considered finished when it is in a terminal condition, either "Complete" or "Failed". A Job cannot have both the "Complete" and "Failed" conditions. Additionally, it cannot be in the "Complete" and "FailureTarget" conditions. The "Complete", "Failed" and "FailureTarget" conditions cannot be disabled.
        /// 
        /// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobCondition> Conditions;
        /// <summary>
        /// The number of pods which reached phase Failed. The value increases monotonically.
        /// </summary>
        public readonly int Failed;
        /// <summary>
        /// FailedIndexes holds the failed indexes when spec.backoffLimitPerIndex is set. The indexes are represented in the text format analogous as for the `completedIndexes` field, ie. they are kept as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the failed indexes are 1, 3, 4, 5 and 7, they are represented as "1,3-5,7". The set of failed indexes cannot overlap with the set of completed indexes.
        /// </summary>
        public readonly string FailedIndexes;
        /// <summary>
        /// The number of active pods which have a Ready condition and are not terminating (without a deletionTimestamp).
        /// </summary>
        public readonly int Ready;
        /// <summary>
        /// Represents time when the job controller started processing a job. When a Job is created in the suspended state, this field is not set until the first time it is resumed. This field is reset every time a Job is resumed from suspension. It is represented in RFC3339 form and is in UTC.
        /// 
        /// Once set, the field can only be removed when the job is suspended. The field cannot be modified while the job is unsuspended or finished.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The number of pods which reached phase Succeeded. The value increases monotonically for a given spec. However, it may decrease in reaction to scale down of elastic indexed jobs.
        /// </summary>
        public readonly int Succeeded;
        /// <summary>
        /// The number of pods which are terminating (in phase Pending or Running and have a deletionTimestamp).
        /// 
        /// This field is beta-level. The job controller populates the field when the feature gate JobPodReplacementPolicy is enabled (enabled by default).
        /// </summary>
        public readonly int Terminating;
        /// <summary>
        /// uncountedTerminatedPods holds the UIDs of Pods that have terminated but the job controller hasn't yet accounted for in the status counters.
        /// 
        /// The job controller creates pods with a finalizer. When a pod terminates (succeeded or failed), the controller does three steps to account for it in the job status:
        /// 
        /// 1. Add the pod UID to the arrays in this field. 2. Remove the pod finalizer. 3. Remove the pod UID from the arrays while increasing the corresponding
        ///     counter.
        /// 
        /// Old jobs might not be tracked using this field, in which case the field remains null. The structure is empty for finished jobs.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Batch.V1.UncountedTerminatedPods UncountedTerminatedPods;

        [OutputConstructor]
        private JobStatus(
            int active,

            string completedIndexes,

            string completionTime,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobCondition> conditions,

            int failed,

            string failedIndexes,

            int ready,

            string startTime,

            int succeeded,

            int terminating,

            Pulumi.Kubernetes.Types.Outputs.Batch.V1.UncountedTerminatedPods uncountedTerminatedPods)
        {
            Active = active;
            CompletedIndexes = completedIndexes;
            CompletionTime = completionTime;
            Conditions = conditions;
            Failed = failed;
            FailedIndexes = failedIndexes;
            Ready = ready;
            StartTime = startTime;
            Succeeded = succeeded;
            Terminating = terminating;
            UncountedTerminatedPods = uncountedTerminatedPods;
        }
    }
}
