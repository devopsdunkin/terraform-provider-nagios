package nagios

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Host contains all info needed to create a host in Nagios
// TODO: Test to see if we need both JSON and schema tags
// Using tag with both JSON and schema because a POST uses URL encoding to send data

// TODO: Need to add in all of the other fields. What we have right now will work for initial testing
type Timeperiod struct {
	TimeperiodName string        `json:"timeperiod_name" schema:"timeperiod_name"`
	Alias          string        `json:"alias"`
	Sunday         string        `json:"sunday,omitempty"`
	Monday         string        `json:"monday,omitempty"`
	Tuesday        string        `json:"tuesday,omitempty"`
	Wednesday      string        `json:"wednesday,omitempty"`
	Thursday       string        `json:"thursday,omitempty"`
	Friday         string        `json:"friday,omitempty"`
	Saturday       string        `json:"saturday,omitempty"`
	Exclude        []interface{} `json:"exclude,omitempty"`
}

/*
	For any bool value, we allow the user to provide a true/false value, but you will notice
	that we immediately convert it to its integer form and then to a string. We want to provide
	the user with an easy to use schema, but Nagios wants the data as a one or zero in string format.
	This seemed to be the easiest way to accomplish that and I wanted to note why it was done that way.
*/

func resourceTimeperiod() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"timeperiod_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the time period",
			},
			"alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"sunday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Determines whether or not the contact will receive notifications about service problems and recoveries",
			},
			"monday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The short name of the time period during which the contact can be notified about host problems or recoveries",
			},
			"tuesday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The short name of the time period during which the contact can be notified about service problems or recoveries",
			},
			"wednesday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host states for which notifications can be sent out to this contact. Valid options are a combination of one or more of the following: d = notify on DOWN host states, u = notify on UNREACHABLE host states, r = notify on host recoveries (UP states), f = notify when the host starts and stops flapping, and s = send notifications",
			},
			"thursday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The service states for which notifications can be sent out to this contact. Valid options are a combination of one or more of the following: w = notify on WARNING service states, u = notify on UNKNOWN service states, c = notify on CRITICAL service states, r = notify on service recoveries (OK states), and f = notify when the service starts and stops flapping.",
			},
			"friday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A list of the short names of the commands used to notify the contact of a host problem or recovery. Multiple notification commands should be separated by commas. All notification commands are executed when the contact needs to be notified",
			},
			"saturday": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A list of the short names of the commands used to notify the contact of a service problem or recovery. Multiple notification commands should be separated by commas. All notification commands are executed when the contact needs to be notified",
			},
			"exclude": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The short name(s) of the contactgroup(s) that the contact belongs to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: resourceCreateTimeperiod,
		Read:   resourceReadTimeperiod,
		Update: resourceUpdateTimeperiod,
		Delete: resourceDeleteTimeperiod,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateTimeperiod(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	timeperiod := setTimeperiodFromSchema(d)

	_, err := nagiosClient.newTimeperiod(timeperiod)

	if err != nil {
		return err
	}

	d.SetId(timeperiod.TimeperiodName)

	return resourceReadTimeperiod(d, m)
}

func resourceReadTimeperiod(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	timeperiod, err := nagiosClient.getTimeperiod(d.Id())

	if err != nil {
		return err
	}

	if timeperiod == nil {
		// contact not found. Let Terraform know to delete the state
		d.SetId("")
		return nil
	}

	setDataFromTimeperiod(d, timeperiod)

	return nil
}

func resourceUpdateTimeperiod(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	timeperiod := setTimeperiodFromSchema(d)

	oldVal, _ := d.GetChange("timeperiod_name")

	if oldVal == "" {
		oldVal = d.Get("timeperiod_name").(string)
	}

	err := nagiosClient.updateTimeperiod(timeperiod, oldVal)

	if err != nil {
		return err
	}

	setDataFromTimeperiod(d, timeperiod)

	return resourceReadTimeperiod(d, m)
}

func resourceDeleteTimeperiod(d *schema.ResourceData, m interface{}) error {
	nagiosClient := m.(*Client)

	_, err := nagiosClient.deleteTimeperiod(d.Id())

	if err != nil {
		return err
	}

	return nil
}

func setDataFromTimeperiod(d *schema.ResourceData, timeperiod *Timeperiod) {
	d.SetId(timeperiod.TimeperiodName)
	d.Set("timeperiod_name", timeperiod.TimeperiodName)
	d.Set("alias", timeperiod.Alias)

	if timeperiod.Sunday != "" {
		d.Set("sunday", timeperiod.Sunday)
	}

	if timeperiod.Monday != "" {
		d.Set("monday", timeperiod.Monday)
	}

	if timeperiod.Tuesday != "" {
		d.Set("tuesday", timeperiod.Tuesday)
	}

	if timeperiod.Wednesday != "" {
		d.Set("wednesday", timeperiod.Wednesday)
	}

	if timeperiod.Thursday != "" {
		d.Set("thursday", timeperiod.Thursday)
	}

	if timeperiod.Friday != "" {
		d.Set("friday", timeperiod.Friday)
	}

	if timeperiod.Saturday != "" {
		d.Set("saturday", timeperiod.Saturday)
	}

	if timeperiod.Exclude != nil {
		d.Set("exclude", timeperiod.Exclude)
	}
}

func setTimeperiodFromSchema(d *schema.ResourceData) *Timeperiod {
	timeperiod := &Timeperiod{
		TimeperiodName: d.Get("timeperiod_name").(string),
		Alias:          d.Get("alias").(string),
		Sunday:         d.Get("sunday").(string),
		Monday:         d.Get("monday").(string),
		Tuesday:        d.Get("tuesday").(string),
		Wednesday:      d.Get("wednesday").(string),
		Thursday:       d.Get("thursday").(string),
		Friday:         d.Get("friday").(string),
		Saturday:       d.Get("saturday").(string),
		Exclude:        d.Get("exclude").(*schema.Set).List(),
	}

	return timeperiod
}
