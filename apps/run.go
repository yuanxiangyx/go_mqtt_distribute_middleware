package apps

import utils2 "mqtt_pro/apps/utils"

func WebApp() error {
	apiRoute := utils2.InitRouter()
	return apiRoute.Run()
}
