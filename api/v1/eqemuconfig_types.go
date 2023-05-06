/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EQEmuConfigSpec defines the desired state of EQEmuConfig
type EQEmuConfigSpec struct {
	ZonesDefaultstatus        string `json:"zones.defaultstatus"`         // "0",
	ZonesPortsLow             string `json:"zones.ports.low"`             // "7000",
	ZonesPortsHigh            string `json:"zones.ports.high"`            // "7400"
	QsdatabaseHost            string `json:"qsdatabase.host"`             // "localhost",
	QsdatabasePort            string `json:"qsdatabase.port"`             // "3306",
	QsdatabaseUsername        string `json:"qsdatabase.username"`         // "root",
	QsdatabasePassword        string `json:"qsdatabase.password"`         // "eqemu",
	QsdatabaseDb              string `json:"qsdatabase.db"`               // "peq"
	ChatserverPort            string `json:"chatserver.port"`             // "7778",
	ChatserverHost            string `json:"chatserver.host"`             // ""
	MailserverHost            string `json:"mailserver.host"`             // "",
	MailserverPort            string `json:"mailserver.port"`             // "7778"
	WorldLoginserver1Account  string `json:"world.loginserver1.account"`  // "",
	WorldLoginserver1Password string `json:"world.loginserver1.password"` // "",
	WorldLoginserver1Legacy   string `json:"world.loginserver1.legacy"`   // "1",
	WorldLoginserver1Host     string `json:"world.loginserver1.host"`     // "login.eqemulator.net",
	WorldLoginserver1Port     string `json:"world.loginserver1.port"`     // "5998"
	WorldLoginserver2Account  string `json:"world.loginserver2.account"`  // "",
	WorldLoginserver2Password string `json:"world.loginserver2.password"` // "",
	WorldLoginserver2Legacy   string `json:"world.loginserver2.legacy"`   // "1",
	WorldLoginserver2Host     string `json:"world.loginserver2.host"`     // "login.eqemulator.net",
	WorldLoginserver2Port     string `json:"world.loginserver2.port"`     // "5998"
	WorldLoginserver3Account  string `json:"world.loginserver3.account"`  // "",
	WorldLoginserver3Password string `json:"world.loginserver3.password"` // "",
	WorldLoginserver3Legacy   string `json:"world.loginserver3.legacy"`   // "1",
	WorldLoginserver3Host     string `json:"world.loginserver3.host"`     // "login.eqemulator.net",
	WorldLoginserver3Port     string `json:"world.loginserver3.port"`     // "5998"
	WorldTcpIp                string `json:"world.tcp.ip"`                // "127.0.0.1",
	WorldTcpPort              string `json:"world.tcp.port"`              // "9001"
	WorldTelnetIp             string `json:"world.telnet.ip"`             // "0.0.0.0",
	WorldTelnetPort           string `json:"world.telnet.port"`           // "9000",
	WorldTelnetEnabled        string `json:"world.telnet.enabled"`        // "true"
	WorldKey                  string `json:"world.key"`                   // "somelongrandomstring12345asdfffff `json:"",
	WorldShortname            string `json:"world.shortname"`             // "Akkas PEQ Installer",
	WorldLongname             string `json:"world.longname"`              // "Akkas PEQ Installer"
	DatabaseDb                string `json:"database.db"`                 // "peq",
	DatabaseHost              string `json:"database.host"`               // "localhost",
	DatabasePort              string `json:"database.port"`               // "3306",
	DatabaseUsername          string `json:"database.username"`           // "root",
	DatabasePassword          string `json:"database.password"`           // "eqemu"
	FilesOpcodes              string `json:"files.opcodes"`               // "assets/opcodes/opcodes.conf",
	FilesMailOpcodes          string `json:"files.mail_opcodes"`          // "assets/opcodes/mail_opcodes.conf"
	DirectoriesPatches        string `json:"directories.patches"`         // "assets/patches/",
	DirectoriesOpcodes        string `json:"directories.opcodes"`         // "assets/opcodes/"
}

// EQEmuConfigStatus defines the observed state of EQEmuConfig
type EQEmuConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EQEmuConfig is the Schema for the eqemuconfigs API
type EQEmuConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EQEmuConfigSpec   `json:"spec,omitempty"`
	Status EQEmuConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EQEmuConfigList contains a list of EQEmuConfig
type EQEmuConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EQEmuConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EQEmuConfig{}, &EQEmuConfigList{})
}
