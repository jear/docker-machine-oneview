package ov

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/docker/machine/log"
	"github.com/docker/machine/drivers/oneview/rest"
)

// firmware
type FirmwareOption struct {
	ForceInstallFirmware  bool    `json:"forceInstallFirmware,omitempty"` // "forceInstallFirmware": false,
	FirmwareBaselineUri   Nstring `json:"firmwareBaselineUri,omitempty"`  // "firmwareBaselineUri": null,
	ManageFirmware        bool    `json:"manageFirmware,omitempty"`       // "manageFirmware": false
}

// Boot mode option
type BootModeOption struct {
	ManageMode     bool       `json:"manageMode,omitempty"`     // "manageMode": true,
	Mode           string     `json:"mode,omitempty"`           // "mode": "BIOS",
	PXEBootPolicy  Nstring    `json:"pxeBootPolicy,omitempty"`  // "pxeBootPolicy": null
}

// Boot management
type BootManagement struct {
	ManageBoot  bool     `json:"manageBoot,omitempty"` // "manageBoot": true,
	Order       []string `json:"order,omitempty"`      // "order": ["CD","USB","HardDisk","PXE"]
}

// Bios Settings
type BiosSettings struct {
	ID    string `json:"id,omitempty"`    // id
	Value string `json:"value,omitempty"` // value
}

// bios options
type BiosOption struct {
	ManageBios          bool `json:"manageBios,omitempty"`                    // "manageBios": false,
	OverriddenSettings []BiosSettings `json:"overriddenSettings,omitempty"`   // "overriddenSettings": []
}

// ServerProfile , server profile object for ov
type ServerProfile struct {
	Affinity               string              `json:"affinity,omitempty"`              // "affinity": "Bay",
	AssociatedServer       Nstring             `json:"associatedServer,omitempty"`      // "associatedServer": null,
	Bios                   BiosOption          `json:"bios,omitempty"`                  // "bios": {	},
	Boot                   BootManagement      `json:"boot,omitempty"`                  // "boot": { },
	BootMode               BootModeOption      `json:"bootMode,omitempty"`              // "bootMode": {},
	Category               string              `json:"category,omitempty"`              // "category": "server-profiles",
	Connections            []Connection        `json:"connections,omitempty"`
	Description            string              `json:"description,omitempty"`           // "description": "Docker Machine Bay 16",
	Created                string              `json:"created,omitempty"`               // "created": "20150831T154835.250Z",
	ETAG                   string              `json:"eTag,omitempty"`                  // "eTag": "1441036118675/8"
	EnclosureBay           int                 `json:"enclosureBay,omitempty"`          // "enclosureBay": 16,
	EnclosureGroupURI      Nstring             `json:"enclosureGroupUri,omitempty"`     // "enclosureGroupUri": "/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039",
	EnclosureURI           Nstring             `json:"enclosureUri,omitempty"`          // "enclosureUri": "/rest/enclosures/092SN51207RR",
	Firmware               FirmwareOption      `json:"firmware,omitempty"`              // "firmware": { },
	HideUnusedFlexNics     bool                `json:"hideUnusedFlexNics,omitempty"`    // "hideUnusedFlexNics": false,
	InProgress             bool                `json:"inProgress,omitempty"`            // "inProgress": false,
	LocalStorage           LocalStorageOptions `json:"localStorage,omitempty"`          // "localStorage": {},
	MACType                string              `json:"macType,omitempty"`               // "macType": "Physical",
	Modified               string              `json:"modified,omitempty"`              // "modified": "20150902T175611.657Z",
	Name                   string              `json:"name,omitempty"`                  // "name": "Server_Profile_scs79",
	SanStorage             SanStorageOptions   `json:"sanStorage,omitempty"`            // "sanStorage": {},
	SerialNumber           Nstring             `json:"serialNumber,omitempty"`          // "serialNumber": "2M25090RMW",
	SerialNumberType       string              `json:"serialNumberType,omitempty"`      // "serialNumberType": "Physical",
	ServerHardwareTypeURI  Nstring             `json:"serverHardwareTypeUri,omitempty"` // "serverHardwareTypeUri": "/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	ServerHardwareURI      Nstring             `json:"serverHardwareUri,omitempty"`     // "serverHardwareUri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57",
	State                  string              `json:"state,omitempty"`                 // "state": "Normal",
	Status                 string              `json:"status,omitempty"`                // "status": "Critical",
	TaskURI                Nstring             `json:"taskUri,omitempty"`               // "taskUri": "/rest/tasks/6F0DF438-7D30-41A2-A36D-62AB866BC7E8",
	Type                   string              `json:"type,omitempty"`                  // 	Type               string `json:"type,omitempty"`	// "type": "ServerProfileV4",
	URI                    Nstring             `json:"uri,omitempty"`                   // "uri": "/rest/server-profiles/9979b3a4-646a-4c3e-bca6-80ca0b403a93",
	UUID                   Nstring             `json:"uuid,omitempty"`                  // "uuid": "30373237-3132-4D32-3235-303930524D57",
	WWNType                string              `json:"wwnType,omitempty"`               // "wwnType": "Physical",
}

// Clone server profile
func (s ServerProfile) Clone() (ServerProfile) {
  var ca []Connection
  for _, c := range s.Connections {
    ca = append(ca, c.Clone())
  }

  return ServerProfile {
				Affinity               : s.Affinity,
				Bios                   : s.Bios,
				Boot                   : s.Boot,
				BootMode               : s.BootMode,
				Connections            : ca,
				Description            : s.Description,
				EnclosureBay           : s.EnclosureBay,
				EnclosureGroupURI      : s.EnclosureGroupURI,
				EnclosureURI           : s.EnclosureURI,
				Firmware               : s.Firmware,
				HideUnusedFlexNics     : s.HideUnusedFlexNics,
				LocalStorage           : s.LocalStorage.Clone(),
				MACType                : s.MACType,
				Name                   : s.Name,
				SanStorage             : s.SanStorage.Clone(),
				SerialNumberType       : s.SerialNumberType,
				Type                   : s.Type,
				WWNType                : s.WWNType,
  }
}

// ServerProfileList a list of ServerProfile objects
// TODO: missing properties, need to think how we can make a higher lvl structure like an OVList
// Then things like Members are inherited
type ServerProfileList struct {
	Total        int      `json:"total,omitempty"`           // "total": 1,
	Count        int      `json:"count,omitempty"`           // "count": 1,
	Start        int      `json:"start,omitempty"`           // "start": 0,
	PrevPageURI  Nstring  `json:"prevPageUri,omitempty"`     // "prevPageUri": null,
	NextPageURI  Nstring  `json:"nextPageUri,omitempty"`     // "nextPageUri": null,
	URI          Nstring  `json:"uri,omitempty"`             // "uri": "/rest/server-profiles?filter=serialNumber%20matches%20%272M25090RMW%27&sort=name:asc"
	Members      []ServerProfile `json:"members,omitempty"`  // "members":[]
}

// get a server profile by name
func (c *OVClient) GetProfileByName(name string)(ServerProfile, error) {
	var (
		profile ServerProfile
	)
	profiles, err := c.GetProfiles( fmt.Sprintf("name matches '%s'",name), "name:asc" )
	if profiles.Total > 0 {
		return profiles.Members[0], err
	} else {
		return profile, err
	}
}

// GetProfileBySN  accepts serial number
func (c *OVClient) GetProfileBySN(serialnum string)(ServerProfile, error) {
	var (
		profile ServerProfile
	)
	profiles, err := c.GetProfiles( fmt.Sprintf("serialNumber matches '%s'",serialnum), "name:asc" )
	if profiles.Total > 0 {
		return profiles.Members[0], err
	} else {
		return profile, err
	}
}

// get a server profiles
func (c *OVClient) GetProfiles(filter string, sort string)(ServerProfileList, error) {
	var (
		uri    = "/rest/server-profiles"
		q      map[string]interface{}
		profiles ServerProfileList
	)
  q = make(map[string]interface{})
	if filter != "" {
		q["filter"] = filter
	}

  if sort != "" {
		q["sort"] = sort
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions( c.GetAuthHeaderMap() )
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri , nil)
	if err != nil {
		return profiles, err
	}

	log.Debugf("GetProfiles %s", data)
	if err := json.Unmarshal([]byte(data), &profiles); err != nil {
		return profiles, err
	}
	return profiles, nil
}

// submit new profile template
func (c *OVClient) SubmitNewProfile(p ServerProfile) (t *Task, err error) {
	log.Infof("Initializing creation of server profile for %s.",p.Name)
	var (
		uri  = "/rest/server-profiles"
	// 	task = rest_api(:oneview, :post, '/rest/server-profiles', { 'body' => new_template_profile })
	)
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, p)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri , p)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new profile request: %s", err)
		return t, err
	}

	log.Debugf("Response NewProfile %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with power state un-marshal: %s", err)
		return t, err
	}

	return t, err
}

// create profile from template
func (c *OVClient) CreateProfileFromTemplate(name string, template ServerProfile, blade ServerHardware) (error) {
	var (
		currenttime int = 0
	)
	log.Infof("TEMPLATE : %+v\n", template)

	var new_template = template.Clone()
	new_template.ServerHardwareURI = blade.URI
	new_template.Description += name
	new_template.Name = name

	t, err := c.SubmitNewProfile(new_template)
	if err != nil { return err }
	for !t.TaskIsDone && (currenttime < t.Timeout) {
		if err := t.GetCurrentTaskStatus(); err != nil {
			return err
		}
		if t.URI != "" && T_COMPLETED.Equal(t.TaskState) {
			t.TaskIsDone = true
		}
		if t.URI != "" {
			log.Debugf("Waiting for task to complete, for %s ", name)
			log.Infof("Working on profile creation,%d%%, %s.", t.ComputedPercentComplete, t.TaskStatus)
		} else {
			log.Info("Working on profile creation.")
		}

		// wait time before next check
		time.Sleep(time.Millisecond * (1000 * t.WaitTime)) // wait 10sec before checking the status again
		currenttime++
	}
	if !(currenttime < t.Timeout) {
		log.Warn("Task timed out.")
	}
	log.Infof("Create server profile Completed")

	// 	60.times do # Wait for up to 5 min
	// 		matching_profiles = rest_api(:oneview, :get, "/rest/server-profiles?filter=name matches '#{host_name}'&sort=name:asc")
	// 		break if matching_profiles['count'] > 0
	// 		print '.'
	// 		sleep 5
	// 	end
	// 	unless matching_profiles['count'] > 0
	// 		task = rest_api(:oneview, :get, task_uri)
	// 		fail "Server profile couldn't be created! #{task['taskStatus']}. #{task['taskErrors'].first['message']}"
	// 	end
	// end
	return nil
}
