package inits

import (
	"path"
	"server/model"
	"server/public"
)

var UserData *model.UsersData

func init()  {
	UserData=model.NewUsersData()
}

func InjectData() error {
	err:=UserData.LoadFromXml(path.Join(public.DATA_FILE_PATH,public.DATA_FILE_NAME))
	if err!=nil{
		return err
	}
	return nil
}
