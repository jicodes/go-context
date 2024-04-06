package main

import (
	"context"
	"errors"
	"fmt"
	"time"

)

func callThirdPartyAPI (ctx context.Context, userID int) (bool, error) {
	
	// Simulate a slow db query
	_ = userID
	time. Sleep (400 * time.Millisecond)

	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("error: context timeout exceeded" )
	}
	return true, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Millisecond)
	defer cancel()

	userId := 4269

	isVerified, err := callThirdPartyAPI(ctx, userId)
	if err != nil {
		println(err.Error())
		return
	}

	if isVerified {
		fmt.Printf("User %v is verified\n", userId)
		return
	}
}