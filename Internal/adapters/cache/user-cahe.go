package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/pkg/text"
)

var (
	user_id_key = "user_id:%d"
	user_username_key = "user_username:%s"
	user_email_key = "user_email:%s"
)

func SetUserCache(ctx context.Context, user model.User) error {
	var (
		keyid = fmt.Sprintf(user_id_key, user.ID)
		keyname = fmt.Sprintf(user_username_key, user.Username)
		keyemail = fmt.Sprintf(user_email_key, user.Email)
	)

	data, err := json.Marshal(user)
	if err != nil {
		return text.ErrEncodeJson
	}

	if err = Set(ctx, keyid, string(data)); err != nil {
		return err
	}
	if err = Set(ctx, keyname, string(data)); err != nil {
		return err
	}
	if err = Set(ctx, keyemail, string(data)); err != nil {
		return err
	}

	return nil
}

func GetUserbyID(ctx context.Context, user *model.User) error {
	var (
		keyid = fmt.Sprintf(user_id_key, user.ID)
	)

	data, err := Get(ctx, keyid)
	if err != nil {
		return err
	}

	var cachedUser model.User
	if err := json.Unmarshal([]byte(data), &cachedUser); err != nil {
		return text.ErrDecodeJson
	}
	*user = cachedUser

	return nil
}

func GetUserbyUsername(ctx context.Context, user *model.User) error {
	var (
		keyname = fmt.Sprintf(user_username_key, user.Username)
	)
	data, err := Get(ctx, keyname)
	if err != nil {
		return err
	}

	var cachedUser model.User
	if err := json.Unmarshal([]byte(data), &cachedUser); err != nil {
		return text.ErrDecodeJson
	}
	*user = cachedUser

	return nil
}

func GetUserbyEmail(ctx context.Context, user *model.User) error {
	var (
		keyemail = fmt.Sprintf(user_email_key, user.Email)
	)
	data, err := Get(ctx, keyemail)
	if err != nil {
		return err
	}

	var cachedUser model.User
	if err := json.Unmarshal([]byte(data), &cachedUser); err != nil {
		return text.ErrDecodeJson
	}
	*user = cachedUser

	return nil
}

func SetUserIDNotfound(ctx context.Context, user model.User) error {
	var (
		keyid = fmt.Sprintf(user_id_key, user.ID)
	)
	if err := SetNotfound(ctx, keyid); err != nil {
		return err
	}
	return nil
}

func SetUserUsernameNotfound(ctx context.Context, user model.User) error {
	var (
		keyname = fmt.Sprintf(user_username_key, user.Username)
	)
	if err := SetNotfound(ctx, keyname); err != nil {
		return err
	}
	return nil
}

func SetUserEmailNotfound(ctx context.Context, user model.User) error {
	var (
		keyemail = fmt.Sprintf(user_email_key, user.Email)
	)
	if err := SetNotfound(ctx, keyemail); err != nil {
		return err
	}
	return nil
}