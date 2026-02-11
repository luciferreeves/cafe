package controllers

import (
	"cafe/models"
	"cafe/repositories"
	"cafe/services/openid"
	"cafe/session"
	"cafe/utils/auth"
	"cafe/utils/shortcuts"
	"context"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	if auth.IsAuthenticated(ctx) {
		return shortcuts.Redirect(ctx, "mainHall")
	}

	state, err := openid.GenerateState()
	if err != nil {
		log.Printf("Failed to generate state: %v", err)
		return InternalServerError(ctx, errors.New("We couldn't start the login process. Please try again."))
	}

	sess, err := session.Store.Get(ctx)
	if err != nil {
		log.Printf("Failed to get session: %v", err)
		return InternalServerError(ctx, errors.New("There was a problem with your session. Please try logging in again."))
	}
	sess.Set("oauth_state", state)
	if err := sess.Save(); err != nil {
		log.Printf("Failed to save session: %v", err)
		return InternalServerError(ctx, errors.New("There was a problem with your session. Please try logging in again."))
	}

	authURL := openid.GetAuthURL(state)
	return ctx.Redirect(authURL)
}

func Callback(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	code := ctx.Query("code")

	if state == "" || code == "" {
		return BadRequest(ctx, errors.New("The login response was incomplete. Please try logging in again."))
	}

	sess, err := session.Store.Get(ctx)
	if err != nil {
		log.Printf("Failed to get session: %v", err)
		return InternalServerError(ctx, errors.New("There was a problem with your session. Please try logging in again."))
	}

	savedState := sess.Get("oauth_state")
	if savedState == nil || savedState.(string) != state {
		return Unauthorized(ctx, errors.New("Your login request could not be verified. Please try again."))
	}

	sess.Delete("oauth_state")

	c := context.Background()
	oauth2Token, err := openid.ExchangeCode(c, code)
	if err != nil {
		log.Printf("Failed to exchange code: %v", err)
		return Unauthorized(ctx, errors.New("We couldn't verify your login with Cloudron. Please try again."))
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		log.Printf("No id_token in response")
		return Unauthorized(ctx, errors.New("Your login was incomplete. Please try again."))
	}

	idToken, err := openid.VerifyIDToken(c, rawIDToken)
	if err != nil {
		log.Printf("Failed to verify ID token: %v", err)
		return Unauthorized(ctx, errors.New("We couldn't verify your identity. Please try logging in again."))
	}

	userInfo, err := openid.GetUserInfo(c, oauth2Token, idToken)
	if err != nil {
		log.Printf("Failed to extract user info: %v", err)
		return InternalServerError(ctx, errors.New("We couldn't retrieve your account information. Please try again."))
	}

	user, err := repositories.GetUserByOpenID(userInfo.Sub)
	isAdminUser := openid.IsAdmin(userInfo)

	if err != nil {
		user = &models.User{
			OpenID:      userInfo.Sub,
			Username:    userInfo.PreferredUsername,
			Email:       userInfo.Email,
			DisplayName: userInfo.Name,
			IsAdmin:     isAdminUser,
		}

		if err := repositories.CreateUser(user); err != nil {
			log.Printf("Failed to create user: %v", err)
			return InternalServerError(ctx, errors.New("We couldn't create your account. Please contact an administrator."))
		}
	} else {
		user.Email = userInfo.Email
		user.DisplayName = userInfo.Name
		user.IsAdmin = isAdminUser
		if err := repositories.UpdateUser(user); err != nil {
			log.Printf("Failed to update user: %v", err)
		}
	}

	sess.Set("username", user.Username)

	if err := sess.Save(); err != nil {
		log.Printf("Failed to save session: %v", err)
		return InternalServerError(ctx, errors.New("We couldn't complete your login. Please try again."))
	}

	return shortcuts.Redirect(ctx, "mainHall")
}

func Logout(ctx *fiber.Ctx) error {
	sess, err := session.Store.Get(ctx)
	if err != nil {
		log.Printf("Failed to get session: %v", err)
		return shortcuts.Redirect(ctx, "mainHall")
	}

	if err := sess.Destroy(); err != nil {
		log.Printf("Failed to destroy session: %v", err)
	}

	return shortcuts.Redirect(ctx, "mainHall")
}
