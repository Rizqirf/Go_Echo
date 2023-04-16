package seeds

import (
	"time"

	"github.com/Rizqirf/go_echo/seed"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
        	{
                        Name: "data 1",
                        Run: func(db *gorm.DB) error {
							return PopulateTable(db, "Promotion 1", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam vel nibh iaculis, sollicitudin diam eu, venenatis metus. Vestibulum eu ligula varius, tincidunt lorem sed, feugiat velit. Sed rutrum diam nec malesuada varius. Ut tempus dui ac vulputate vestibulum. Donec commodo sem nulla, non auctor quam lacinia eu. Etiam eu quam laoreet, tristique neque luctus, vehicula ligula. Vestibulum placerat magna vitae efficitur finibus. Maecenas eu volutpat lectus. Nam euismod libero eu tincidunt tincidunt. Aliquam vitae fringilla leo. Aliquam efficitur congue risus eu pretium. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae;","https://picsum.photos/700/200", datatypes.Date(time.Now()),datatypes.Date(time.Now()))
                        },
                },
        	{
                        Name: "data 2",
                        Run: func(db *gorm.DB) error {
							return PopulateTable(db, "Promotion 2", "Quisque rutrum mollis libero, eget blandit felis dapibus non. Suspendisse vitae accumsan nunc, ac fringilla nunc. Suspendisse arcu eros, sollicitudin eget purus eget, molestie laoreet libero. Integer vulputate urna non elit tincidunt, ac fermentum justo ultrices. Sed risus dolor, sodales sed convallis in, fermentum sit amet nunc. Ut condimentum nec turpis sed molestie. Praesent rutrum consectetur sapien. Curabitur lobortis facilisis luctus. Fusce vehicula mauris id lectus posuere, non sodales sapien rutrum. Pellentesque ac libero dignissim, vestibulum sapien sit amet, iaculis metus. Morbi vel congue mauris, quis ultrices ex.","https://picsum.photos/700/200", datatypes.Date(time.Now()),datatypes.Date(time.Now()))
                        },
                },
        }
}