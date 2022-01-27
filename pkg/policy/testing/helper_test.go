package testing

import (
	"log"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/policy"
)

func TestManager_Load(t *testing.T) {
	// Skip test for now since github is annoying
	files, err := FilePathWalkDir("../../../database-data")
	log.Println(files, err)

}

func TestManager_Group(t *testing.T) {
	// Skip test for now since github is annoying
	files, _ := FilePathWalkDir("../../../database-data")
	log.Printf("%+v", FilterFiles(files))

}

func TestManager_Execute(t *testing.T) {
	// Skip test for now since github is annoying
	TestPolicy(t, policy.Policy{
		Views: []*policy.View{
			{
				Name: "aws_security_group_ingress_rules",
				Query: `WITH sg_rules_ports AS (
					SELECT
						sg.account_id,
						sg.region,
						sg.group_name,
						sg.arn,
						sg.id,
						p.from_port,
						p.to_port,
						p.ip_protocol,
						p.cq_id AS permission_id
					FROM aws_ec2_security_groups sg
						LEFT JOIN
							aws_ec2_security_group_ip_permissions p ON
								sg.cq_id = p.security_group_cq_id
					WHERE p.permission_type = 'ingress'
				)
				SELECT sgs.*,
					r.cidr AS ip
				FROM sg_rules_ports sgs
					LEFT JOIN
						aws_ec2_security_group_ip_permission_ip_ranges r ON
							sgs.permission_id = r.security_group_ip_permission_cq_id
				`,
			},
		},
		Checks: []*policy.Check{
			{
				Name: "Test",
				Query: `SELECT account_id,
				region,
				group_name,
				id,
				from_port,
				to_port,
				ip_protocol,
				ip
		 FROM aws_security_group_ingress_rules
		 WHERE (ip = '0.0.0.0/0' OR ip = '::/0')
			 AND (from_port IS NULL AND to_port IS NULL) -- all ports
			 OR from_port IS DISTINCT FROM 80
			 OR to_port IS DISTINCT FROM 80
			 OR from_port IS DISTINCT FROM 443
			 OR to_port IS DISTINCT FROM 443;`,
			},
		},
	})

}
