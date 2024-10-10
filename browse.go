package main

import (
	"bloggy/internal/database"
	"context"
	"errors"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit, err := setLimit(cmd.args)
	if err != nil {
		return err
	}

	ctx := context.Background()
	gpfuParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	userPosts, err := s.db.GetPostsForUser(ctx, gpfuParams)
	if err != nil {
		return err
	}

	for _, post := range userPosts {
		fmt.Printf("\nTitle: %v \nURL: %v\n", post.Title, post.Url)
	}

	return nil
}

func setLimit(args []string) (int, error) {
	if len(args) > 1 {
		return 0, errors.New("Browse takes only one additional parameter for number of feeds")
	}

	if len(args) == 0 {
		return 2, nil
	} else {
		limit, err := strconv.Atoi(args[0])
		if err != nil {
			return 0, err
		}
		return limit, nil
	}
}
