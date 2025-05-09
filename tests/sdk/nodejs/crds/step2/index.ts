// Copyright 2016-2019, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as pulumi from "@pulumi/pulumi";
import * as k8s from "@pulumi/kubernetes";

const namespace = new k8s.core.v1.Namespace("test-namespace");

//
// Create a CustomResourceDefinition and a CustomResource.
//

const crd = new k8s.apiextensions.v1.CustomResourceDefinition("foobar", {
    metadata: { name: "foobars.stable.example.com" },
    spec: {
        group: "stable.example.com",
        versions: [
            {
                name: "v1",
                served: true,
                storage: true,
                schema: {
                    openAPIV3Schema: {
                        type: "object",
                        properties: {
                            spec: {
                                type: "object",
                                properties: {
                                    foo: {
                                        type: "string",
                                    },
                                    bar: {
                                        type: "string",
                                    },
                                },
                            }
                        }
                    }
                }
            }
        ],
        scope: "Namespaced",
        names: {
            plural: "foobars",
            singular: "foobar",
            kind: "FooBar",
            shortNames: ["fb"]
        }
    }
});

export const preserveUnknownFields = crd.spec.versions[0].schema.openAPIV3Schema.x_kubernetes_preserve_unknown_fields

new k8s.apiextensions.CustomResource(
    "my-new-foobar-object",
    {
        apiVersion: "stable.example.com/v1",
        kind: "FooBar",
        metadata: {
            namespace: namespace.metadata.name,
            name: "my-new-foobar-object",
        },
        spec: { foo: "such amaze", bar: "wow" }
    },
);

//
// Get the CustomResource.
//
export const objRef = k8s.apiextensions.CustomResource.get("my-new-foobar-object-ref", {
    apiVersion: "stable.example.com/v1",
    kind: "FooBar",
    id: pulumi.interpolate`${namespace.metadata.name}/my-new-foobar-object`
});
