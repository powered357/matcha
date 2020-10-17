package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"strconv"
	"net/http"
)

// HTTP HANDLER FOR DOMAIN /message/get/ . IT HANDLES:
// IT RETURNS OWN USER DATA IN RESPONSE BY POST METHOD.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) LikeUnset(w http.ResponseWriter, r *http.Request) {
	var (
		myUid, otherUid int
		requestParams	map[string]interface{}
		item			interface{}
		ok, isExist     bool
		err  error
		ctx  context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)
	requestParams, ok = ctx.Value("requestParams").(map[string]interface{})
	if !ok {
		server.Logger.LogWarning(r, "Request params has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Параметры запроса имеют неверный тип", "request params has wrong type"))
		return
	}
	item, isExist = requestParams["otherUid"]
	if !isExist {
		server.Logger.LogWarning(r, "Param otherUid expected")
		server.error(w, errors.NoArgument.WithArguments("отсутствует параметр otherUid", "param otherUid expected"))
		return
	}
	otherUid, ok = item.(int)
	if !ok {
		server.Logger.LogWarning(r, "Id of another user has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("otherUid имеет неверный тип", "otherUid has wrong type"))
		return
	}

	err = server.Db.UnsetLike(myUid, otherUid)
	if errors.ImpossibleToExecute.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "Imposible to set like from user#"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
			" to user#"+BLUE+strconv.Itoa(otherUid)+NO_COLOR)
		server.error(w, errors.ImpossibleToExecute.WithArguments("Лайк отсутствует в базе", "like not exists in database"))
		return
	} else if err != nil {
		server.Logger.LogError(r, "SetNewLike returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "Like was unset successfully from user #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		" to user #"+BLUE+strconv.Itoa(otherUid)+NO_COLOR)
}