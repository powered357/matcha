package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /notification/get/ . IT HANDLES:
// IT RETURNS OWN USER DATA IN RESPONSE BY POST METHOD.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) NotificationGet(w http.ResponseWriter, r *http.Request) {
	var (
		notifs []Notif
		myUid  int
		err    error
		ctx    context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)

	notifs, err = server.Db.GetNotifByUidReceiver(myUid)
	if err != nil {
		server.Logger.LogError(r, "GetNotifByUidReceiver returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	jsonNotifs, err := json.Marshal(notifs)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonNotifs)
	server.Logger.LogSuccess(r, "Notifications was handled successfully. Amount is #"+BLUE+strconv.Itoa(len(notifs))+NO_COLOR)
}
