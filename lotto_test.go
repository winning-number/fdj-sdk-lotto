package lotto

/*
// test data
const (
	folderPathCSV0 = "folder_0"
	folderPathCSV1 = "folder_1"
	folderPathCSV2 = "folder_2"
	folderPathCSV3 = "folder_3"
	folderPathCSV4 = "folder_4"
)

var DataCSV0 = DrawCSV0{
	Date:           "20200219",
	ForclosureDate: "20200419",
	Day:            "VENDREDI",
}

var DataCSV1 = DrawCSV1{
	Date:           "20200219",
	ForclosureDate: "20200419",
	Day:            "VENDREDI",
}

var DataCSV2 = DrawCSV2{
	Date:           "2020/02/19",
	ForclosureDate: "2020/04/19",
	Day:            "VENDREDI",
}

var DataCSV3 = DrawCSV3{
	Date:           "2020/02/19",
	ForclosureDate: "2020/04/19",
	Day:            "VENDREDI",
}

var DataCSV4 = DrawCSV4{
	Date:           "2020/02/19",
	ForclosureDate: "2020/04/19",
	Day:            "VENDREDI",
}

func TestNew(t *testing.T) {
	expectedTimeErrNoSeparator := "parsing time \"\" as \"20060102\": cannot parse \"\" as \"2006\""
	expectedTimeErrSlachSeparator := "parsing time \"\" as \"2006/01/02\": cannot parse \"\" as \"2006\""
	t.Run("Should be ok", func(t *testing.T) {
		l, err := New(Folders{
			CSV0FolderPath: folderPathCSV0,
			CSV1FolderPath: folderPathCSV1,
			CSV2FolderPath: folderPathCSV2,
			CSV3FolderPath: folderPathCSV3,
			CSV4FolderPath: folderPathCSV4,
		})

		driver, ok := l.(*lotto)
		if !ok {
			t.Fatal("New return bad lotto type")
		}
		if assert.NoError(t, err) {
			assert.NotNil(t, driver.parser)
			assert.Empty(t, driver.draws)

			/*
			** Test configurations
*/ /*
			if assert.Len(t, driver.confs, 5) {
				// test config for the DrawCSV0 model
				assert.Equal(t, folderPathCSV0, driver.confs[0].FolderPath)
				if assert.NotNil(t, driver.confs[0].CreateObject) {
					assert.IsType(t, &DrawCSV0{}, driver.confs[0].CreateObject())
					assert.NoError(t, driver.confs[0].RecordObject(&DataCSV0, ""))
					assert.EqualError(t, driver.confs[0].RecordObject(&DrawCSV0{}, ""), expectedTimeErrNoSeparator)
					assert.EqualError(t, driver.confs[0].RecordObject(&Metadata{}, ""), ErrDrawTypeDecode.Error())
				}

				// test config for the DrawCSV1 model
				assert.Equal(t, folderPathCSV1, driver.confs[1].FolderPath)
				if assert.NotNil(t, driver.confs[1].CreateObject) {
					assert.IsType(t, &DrawCSV1{}, driver.confs[1].CreateObject())
					assert.NoError(t, driver.confs[1].RecordObject(&DataCSV1, ""))
					assert.EqualError(t, driver.confs[1].RecordObject(&DrawCSV1{}, ""), expectedTimeErrNoSeparator)
					assert.EqualError(t, driver.confs[1].RecordObject(&Metadata{}, ""), ErrDrawTypeDecode.Error())
				}

				// test config for the DrawCSV2 model
				assert.Equal(t, folderPathCSV2, driver.confs[2].FolderPath)
				if assert.NotNil(t, driver.confs[2].CreateObject) {
					assert.IsType(t, &DrawCSV2{}, driver.confs[2].CreateObject())
					assert.NoError(t, driver.confs[2].RecordObject(&DataCSV2, ""))
					assert.EqualError(t, driver.confs[2].RecordObject(&DrawCSV2{}, ""), expectedTimeErrSlachSeparator)
					assert.EqualError(t, driver.confs[2].RecordObject(&Metadata{}, ""), ErrDrawTypeDecode.Error())
				}

				// test config for the DrawCSV3 model
				assert.Equal(t, folderPathCSV3, driver.confs[3].FolderPath)
				if assert.NotNil(t, driver.confs[3].CreateObject) {
					assert.IsType(t, &DrawCSV3{}, driver.confs[3].CreateObject())
					assert.NoError(t, driver.confs[3].RecordObject(&DataCSV3, ""))
					assert.EqualError(t, driver.confs[3].RecordObject(&DrawCSV3{}, ""), expectedTimeErrSlachSeparator)
					assert.EqualError(t, driver.confs[3].RecordObject(&Metadata{}, ""), ErrDrawTypeDecode.Error())
				}

				// test config for the DrawCSV4 model
				assert.Equal(t, folderPathCSV4, driver.confs[4].FolderPath)
				if assert.NotNil(t, driver.confs[4].CreateObject) {
					assert.IsType(t, &DrawCSV4{}, driver.confs[4].CreateObject())
					assert.NoError(t, driver.confs[4].RecordObject(&DataCSV4, ""))
					assert.EqualError(t, driver.confs[4].RecordObject(&DrawCSV4{}, ""), expectedTimeErrSlachSeparator)
					assert.EqualError(t, driver.confs[4].RecordObject(&Metadata{}, ""), ErrDrawTypeDecode.Error())
				}

				// final evaluation
				assert.Len(t, driver.draws, 5)
			}
		}
	})
}
*/
