package port

import "github.com/sifer169966/device-interactions/internal/core/domain"

/*
	|--------------------------------------------------------------------------
	| Application Port
	|--------------------------------------------------------------------------
	|
	| Here you can define an interface which will allow foreign actors to
	| communicate with the Application
	|
*/

type DeviceInteractionsRepository interface {
	CreateMany(d domain.DeviceSubmission) error
}
