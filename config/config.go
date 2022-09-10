package config

func InitAllConfigs() (err error) {
	err = InitOSSConfig()
	if err != nil {
		return err
	}
	return nil
}
