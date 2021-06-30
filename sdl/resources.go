package sdl

import (
	"github.com/ovrclk/akash/types"
)

type v2ComputeResources struct {
	CPU     *v2ResourceCPU         `yaml:"cpu"`
	Memory  *v2ResourceMemory      `yaml:"memory"`
	Storage v2ResourceStorageArray `yaml:"storage"`
}

func (sdl *v2ComputeResources) toResourceUnits() types.ResourceUnits {
	if sdl == nil {
		return types.ResourceUnits{}
	}

	var units types.ResourceUnits
	if sdl.CPU != nil {
		units.CPU = &types.CPU{
			Units:      types.NewResourceValue(uint64(sdl.CPU.Units)),
			Attributes: types.Attributes(sdl.CPU.Attributes),
		}
	}
	if sdl.Memory != nil {
		units.Memory = &types.Memory{
			Quantity:   types.NewResourceValue(uint64(sdl.Memory.Quantity)),
			Attributes: types.Attributes(sdl.Memory.Attributes),
		}
	}

	for _, storage := range sdl.Storage {
		storageEntry := types.Storage{
			Quantity:   types.NewResourceValue(uint64(storage.Quantity)),
			Attributes: types.Attributes(storage.Attributes),
		}

		units.Storage = append(units.Storage, storageEntry)
	}

	return units
}
