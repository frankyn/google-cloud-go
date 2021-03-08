// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recommender_test

import (
	"context"

	recommender "cloud.google.com/go/recommender/apiv1beta1"
	"google.golang.org/api/iterator"
	recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_ListInsights() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.ListInsightsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListInsights(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetInsight() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.GetInsightRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetInsight(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkInsightAccepted() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.MarkInsightAcceptedRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MarkInsightAccepted(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListRecommendations() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.ListRecommendationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListRecommendations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetRecommendation() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.GetRecommendationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetRecommendation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkRecommendationClaimed() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.MarkRecommendationClaimedRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MarkRecommendationClaimed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkRecommendationSucceeded() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.MarkRecommendationSucceededRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MarkRecommendationSucceeded(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkRecommendationFailed() {
	// import recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"

	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommenderpb.MarkRecommendationFailedRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MarkRecommendationFailed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
