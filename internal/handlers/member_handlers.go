package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dlsu-lscs/lscs-central-auth-api/internal/database"
	"github.com/dlsu-lscs/lscs-central-auth-api/internal/repository"
	"github.com/labstack/echo/v4"
)

type EmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func GetMemberInfo(c echo.Context) error {
	ctx := c.Request().Context()
	dbconn := database.Connect()
	q := repository.New(dbconn)

	req := new(EmailRequest)
	if err := c.Bind(req); err != nil {
		log.Printf("Failed to parse request body: %v", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	// TODO: add go-validator for validating request body to structs
	// if err := c.Validate(req); err != nil {
	// 	log.Printf("Validation error: %v", err)
	// 	return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid email format"})
	// }

	memberInfo, err := q.GetMemberInfo(ctx, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Email is not an LSCS member"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"email":          memberInfo.Email,
		"full_name":      memberInfo.FullName,
		"committee_name": memberInfo.CommitteeName,
		"division_name":  memberInfo.DivisionName,
		"position_name":  memberInfo.PositionName,
	})
}

func GetAllMembersHandler(c echo.Context) error {
	ctx := c.Request().Context()
	dbconn := database.Connect()
	queries := repository.New(dbconn)

	members, err := queries.ListMembers(ctx)
	if err != nil {
		log.Printf("Failed to list members: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to list members"})
	}

	return c.JSON(http.StatusOK, members)
}

// this will be included in the google auth callback handler
func CheckEmailHandler(c echo.Context) error {
	req := new(EmailRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email is required"})
	}

	ctx := c.Request().Context()
	dbconn := database.Connect()
	queries := repository.New(dbconn)
	memberEmail, err := queries.CheckEmailIfMember(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "Not an LSCS member",
				"state": "absent",
				"email": memberEmail,
			})
		}
		log.Printf("Error checking email: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": "Email is an LSCS member",
		"state":   "present",
		"email":   memberEmail,
	})
}

func RefreshTokenHandler(c echo.Context) error {
	// get refresh token from request header frfr
	// get hashed token from database
	// call CompareTokens to compare
	// if valid, tokens.GenerateJWT (generate new access token)
	// --> also generate new refreshToken maybe (call tokens.GenerateRefreshToken)
	// --> then store newRefreshToken in the database
	return c.JSON(http.StatusOK, echo.Map{
		"access_token": "return new access token here", // TODO: handle refreshing tokens
		// "refresh_token": newRefreshToken,
	})
}

func GetAllCommitteesHandler(c echo.Context) error {
	ctx := c.Request().Context()
	dbconn := database.Connect()
	queries := repository.New(dbconn)
	committees, err := queries.GetAllCommittees(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("Internal server error: %v", err),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"committees": committees,
	})
}
