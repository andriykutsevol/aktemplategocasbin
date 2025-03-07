package application

import (
	"context"
	"log"
	"os"

	"github.com/andriykutsevol/WeatherServer/internal/domain/menu"
	"github.com/andriykutsevol/WeatherServer/internal/domain/menu/menuaction"
	"github.com/andriykutsevol/WeatherServer/internal/domain/menu/menuactionresource"

	"github.com/andriykutsevol/WeatherServer/internal/domain/user/role"
	"github.com/andriykutsevol/WeatherServer/internal/domain/user/rolemenu"

	"github.com/andriykutsevol/WeatherServer/internal/domain/user"
	"github.com/andriykutsevol/WeatherServer/internal/domain/user/userrole"

	"github.com/andriykutsevol/WeatherServer/internal/domain/weather"

	//"github.com/andriykutsevol/WeatherServer/internal/domain/pagination"
	//"github.com/andriykutsevol/WeatherServer/internal/domain/trans"

	"github.com/andriykutsevol/WeatherServer/pkg/util/hash"
	"github.com/andriykutsevol/WeatherServer/pkg/util/uuid"
	"github.com/andriykutsevol/WeatherServer/pkg/util/yaml"
)

type Seed interface {
	Execute(ctx context.Context) error
}

type SeedMenus []struct {
	Name     string `yaml:"name"`
	Icon     string `yaml:"icon"`
	Router   string `yaml:"router,omitempty"`
	Sequence int    `yaml:"sequence"`
	Actions  []struct {
		Code      string `yaml:"code"`
		Name      string `yaml:"name"`
		Resources []struct {
			Method string `yaml:"method"`
			Path   string `yaml:"path"`
		} `yaml:"resources"`
	} `yaml:"actions,omitempty"`
	Children SeedMenus
}


type SeedRoles []struct {
	Name     string `yaml:"name"`
	Sequence int    `yaml:"sequence"`
	Memo string		`yaml:"memo"`
	Rolemenus  []struct {
		MenuID	string	 `yaml:"menuid"`
		ActionID string  `yaml:"actionid"`
	} `yaml:"rolemenus,omitempty"`
}



type SeedUsers []struct {
	Name     string `yaml:"name"`
	RealName     string `yaml:"realname"`
	Password     string `yaml:"password"`
	Email     string `yaml:"email"`
	Phone     string `yaml:"phone"`
    RoleIDs  []struct {
        RoleID string `yaml:"roleid"`
    } `yaml:"roleids"`
}



func NewSeed(
	menuSvc menu.Service,
	menuactionrepo menuaction.Repository,
	roleRepo role.Repository,
	roleMenuRepo rolemenu.Repository,
	userRepo user.Repository,
	userRoleRepo userrole.Repository,
	weatherRepo weather.Repository,
) Seed {
	return &seedApp{
		menuSvc:   menuSvc,
		menuactionrepo: menuactionrepo,
		roleRepo:  roleRepo,
		roleMenuRepo: roleMenuRepo,
		userRepo: userRepo,
		userRoleRepo: userRoleRepo,
		weatherRepo: weatherRepo,
	}
}

type seedApp struct {
	menuSvc      menu.Service
	menuactionrepo menuaction.Repository
	roleRepo     role.Repository
	roleMenuRepo rolemenu.Repository
	userRepo user.Repository
	userRoleRepo userrole.Repository
	weatherRepo weather.Repository
}


//================================================================================================
//================================================================================================

// If we run this, we suppose that db schema is empty
// In the development mode we just purge all tables.
func (s seedApp) Execute(ctx context.Context) error {

	if err := s.menuSeed(ctx, "../../configs/menu.yaml"); err != nil {
		return err
	}
	if err := s.roleSeed(ctx, "../../configs/role.yaml"); err != nil {
		return err
	}
	if err := s.userSeed(ctx, "../../configs/user.yaml"); err != nil {
		return err
	}
	
	// if err := s.weatherSeed(ctx); err != nil {
	// 	return err
	// }
	log.Println("Seed database has been created")

	return nil
}


//================================================================================================
//================================================================================================



func (s seedApp) userSeed(ctx context.Context, menuRolePath string) error {

	err := s.userRepo.Purge(ctx)
	if err != nil{
		return err
	}
	
	err = s.userRoleRepo.Purge(ctx)
	if err != nil{
		return err
	}

	data, err := s.readUserData(menuRolePath)
	if err != nil {
		return err
	}
	err = s.createUsers(ctx, data)
	if err != nil {
		return err
	}

	return nil
}



func (s seedApp) createUsers(ctx context.Context, usersSeed SeedUsers) error {

	for _, userSeed:= range usersSeed{

		user := &user.User{
			ID: uuid.MustString(),
			UserName: userSeed.Name,
			RealName: userSeed.RealName,
			Password: hash.SHA1String(userSeed.Password),
			Email: &userSeed.Email,
			Phone: &userSeed.Phone,
			Status: 1,
			IDString: &userSeed.Name,
		}

		var userRoles userrole.UserRoles

		for _, userRoleID := range userSeed.RoleIDs {

			roleIDString := userRoleID.RoleID
			idString := *user.IDString + "::" + roleIDString
			userIDString := user.IDString

			userRoles = append(userRoles, &userrole.UserRole{
				ID: 		uuid.MustString(),
				UserID: 	user.ID,
				RoleID: 	userRoleID.RoleID,
				IDString: 	&idString,
				UserIDString: userIDString,
				RoleIDString: &roleIDString,
			})
		}

		err := s.createUser(ctx, user, &userRoles)
		if err != nil {
			return err
		}	
	}
	return nil
}




func (s seedApp) createUser(ctx context.Context, user *user.User, userRoles *userrole.UserRoles) (error) {

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	for _ , userrole := range *userRoles {

		role, err := s.roleRepo.GetByIdString(ctx, *userrole.RoleIDString)
		if err != nil {
			return err
		}

		userrole.RoleID = role.ID

		err = s.userRoleRepo.Create(ctx, userrole)
		if err != nil {
			return err
		}
	}
	return nil
}



func (s seedApp) readUserData(name string) (SeedUsers, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data SeedUsers
	d := yaml.NewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err
}


//================================================================================================
//================================================================================================

func (s seedApp) roleSeed(ctx context.Context, menuRolePath string) error {
	s.roleRepo.Purge(ctx)
	s.roleMenuRepo.Purge(ctx)

	data, err := s.readRoleData(menuRolePath)
	if err != nil {
		return err
	}

	err = s.createRoles(ctx, data)
	if err != nil {
		return err
	}		

	return nil
}


func (s seedApp) createRoles(ctx context.Context, list SeedRoles) error {
	for _, item:= range list{

		roleid := uuid.MustString()
		var rms rolemenu.RoleMenus
		
		for _, rmenu := range item.Rolemenus{
			
			// rmenu is a loop variable, and its address does not change across iterations.
			// Loop variables like rmenu in a for loop are reused for each iteration, 
			// meaning the address of rmenu remains the same throughout the loop. 
			// This can cause issues when taking the address of a loop variable, 
			// leading to all elements pointing to the same address.


			// By creating new variables inside the loop and taking their addresses, 
			// you ensure that each RoleMenu object has unique pointers for RoleIDString, MenuIDString, and ActionIDString
			roleIDString := item.Name
			menuIDString := rmenu.MenuID
			actionIDString := rmenu.ActionID

			rm := &rolemenu.RoleMenu{
				RoleID:         roleid,
				MenuID:         rmenu.MenuID,
				ActionID:       rmenu.ActionID,
				RoleIDString:   &roleIDString,
				MenuIDString:   &menuIDString,
				ActionIDString: &actionIDString,
			}
			// Another approach to do that.
			// rm := &rolemenu.RoleMenu{
			// 	RoleID: roleid,
			// 	MenuID: rmenu.MenuID,
			// 	ActionID: rmenu.ActionID,
			// 	RoleIDString: func(s string) *string { return &s }(item.Name),
			// 	MenuIDString: func(s string) *string { return &s }(rmenu.MenuID),
			// 	ActionIDString: func(s string) *string { return &s }(rmenu.ActionID),
			// }

			rms = append(rms, rm)
		}

		ritem := &role.Role{
			ID: roleid,
			Name: item.Name,
			Sequence: item.Sequence,
			Memo: &item.Memo,
			Status: 1,
			RoleMenus: rms,
			IDString: &item.Name,
		}

		err := s.createRole(ctx, ritem)
		if err != nil {
			return err
		}		
	}
	return nil
}




func (s seedApp) createRole(ctx context.Context, item *role.Role) (error) {
	err := s.roleRepo.Create(ctx, item)
	if err != nil {
		return err
	}

	for _ , rmItem := range item.RoleMenus {

		rmItem.ID = uuid.MustString()
		rmItem.RoleID = item.ID

		menu, err := s.menuSvc.GetByIdString(ctx, *rmItem.MenuIDString)
		if err != nil {
			return err
		}
		rmItem.MenuID = menu.ID

		menuAction, err := s.menuactionrepo.GetByIdString(ctx, *rmItem.ActionIDString)
		if err != nil {
			return err
		}
		rmItem.ActionID = menuAction.ID

		rmItem.IDString = new(string)
		*rmItem.IDString = *item.IDString + "::" + *rmItem.ActionIDString

		rmItem.RoleIDString = item.IDString

		err = s.roleMenuRepo.Create(ctx, rmItem)
		if err != nil {
			return err
		}
	}
	return nil
}




func (s seedApp) readRoleData(name string) (SeedRoles, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data SeedRoles
	d := yaml.NewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err

}


//================================================================================================
//================================================================================================


func (s seedApp) menuSeed(ctx context.Context, menuSeedPath string) error {
	// TODO. We are not doing any checks at this time. We assume that menu will be created from scratch.
	
	s.roleMenuRepo.Purge(ctx)
	s.userRoleRepo.Purge(ctx)
	s.menuSvc.PurgeMmenu(ctx)

	data, err := s.readMenuData(menuSeedPath)
	if err != nil {
		return err
	}	
	err = s.createMenus(ctx, "", data)
	if err != nil {
		return err
	}
	return nil 

}


func (s seedApp) readMenuData(name string) (SeedMenus, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data SeedMenus
	d := yaml.NewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	//log.Printf("%+v", data)
	return data, err
}

func (s seedApp) createMenus(ctx context.Context, parentID string, list SeedMenus) error {
	for _, item := range list {
		var as menuaction.MenuActions
		for _, action := range item.Actions {
			var ars menuactionresource.MenuActionResources
			for _, r := range action.Resources {
				ars = append(ars, &menuactionresource.MenuActionResource{
					Method: r.Method,
					Path:   r.Path,
				})
			}
			as = append(as, &menuaction.MenuAction{
				Code:      action.Code,
				Name:      action.Name,
				Resources: ars,
			})
		}
		sitem := &menu.Menu{
			Name:       item.Name,
			Sequence:   item.Sequence,
			Icon:       item.Icon,
			Router:     item.Router,
			ParentID:   parentID,
			Status:     1,
			ShowStatus: 1,
			Actions:    as,
		}

		menuID, err := s.menuSvc.Create(ctx, sitem)
		if err != nil {
			return err
		}

		if item.Children != nil && len(item.Children) > 0 {
			err := s.createMenus(ctx, menuID, item.Children)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


//================================================================================================
//================================================================================================


func (s seedApp) weatherSeed(ctx context.Context) error {

    seedParams := map[string]string{
        "id": "dnipro",
    }
	s.weatherRepo.Seed(ctx, seedParams)

	return nil
}


