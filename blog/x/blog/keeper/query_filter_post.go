package keeper

import (
	"context"

	"blog/x/blog/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"strings"
)

func (k Keeper) FilterPost(ctx context.Context, req *types.QueryFilterPostRequest) (*types.QueryFilterPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	var keyword = req.Keyword

	var posts []types.Post
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var post types.Post
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}

		if strings.Contains(post.Title, keyword) {
			posts = append(posts, post)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFilterPostResponse{Post: posts, Pagination: pageRes}, nil
}
