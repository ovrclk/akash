package sdl

import (
	"sort"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/ovrclk/akash/types"
)

const (
	storageAttributePersistent = "persistent"
	storageAttributeClass      = "class"
	storageAttributeMount      = "mount"
	storageAttributeReadOnly   = "readOnly" // we might not need it at this point of time
)

var (
	errUnsupportedStorageAttribute = errors.New("sdl: unsupported storage attribute")
	errStorageMountPoint           = errors.New("sdl: persistent storage must have mount point")
	errStorageMultipleEphemeral    = errors.New("sdl: multiple ephemeral storages are not allowed")
)

type v2StorageAttributes types.Attributes

type v2ResourceStorage struct {
	Quantity   byteQuantity        `yaml:"size"`
	Attributes v2StorageAttributes `yaml:"attributes,omitempty"`
}

type v2ResourceStorageArray []v2ResourceStorage

var allowedStorageAttributes = map[string]bool{
	storageAttributePersistent: true,
	storageAttributeClass:      true,
	storageAttributeMount:      true,
	storageAttributeReadOnly:   true,
}

// UnmarshalYAML unmarshal storage config
// data can be present either as single entry mapping or an array of them
// e.g
// single entity
// ```yaml
// storage:
//   quantity: 1Gi
//   attributes:
//     class: ssd
// ```
//
// ```yaml
// storage:
//   - quantity: 512Mi # ephemeral storage
//   - quantity: 1Gi
//     attributes:
//       class: ssd
//   - quantity: 100Gi
//     attributes:
//       persistent: true # this volumes survives pod restart
//       class: gp # aka general purpose
// ```
func (sdl *v2ResourceStorageArray) UnmarshalYAML(node *yaml.Node) error {
	var nodes v2ResourceStorageArray

	switch node.Kind {
	case yaml.SequenceNode:
		for _, content := range node.Content {
			var nd v2ResourceStorage
			if err := content.Decode(&nd); err != nil {
				return err
			}

			nodes = append(nodes, nd)
		}
	case yaml.MappingNode:
		var nd v2ResourceStorage
		if err := node.Decode(&nd); err != nil {
			return err
		}

		nodes = append(nodes, nd)
	}

	// check storage parameters in SDL are set properly
	// 1. there can be only one ephemeral entry
	//    ephemeral considered as one without mount point
	// 2. ephemeral cannot have persistence
	// 3. ephemeral cannot be readOnly (for now. configMaps are considered as ephemeral as well)
	// todo make this as helper
	ephemeral := false
	for _, nd := range nodes {
		attr := make(map[string]string)

		for _, nd := range nd.Attributes {
			attr[nd.Key] = nd.Value
		}

		if _, persistence := attr[storageAttributePersistent]; persistence && len(attr[storageAttributeMount]) == 0 {
			return errStorageMountPoint
		}

		if len(attr[storageAttributeMount]) == 0 && ephemeral {
			return errStorageMultipleEphemeral
		} else {
			ephemeral = true
		}
	}

	// sort storage slice in the following order
	// 1. smaller size
	// 2. if sizes are equal then one without mount point goes up
	// 3. if both mount points present use lexical order
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Quantity < nodes[j].Quantity {
			return true
		}

		if nodes[i].Quantity > nodes[j].Quantity {
			return false
		}

		iAttr := types.Attributes(nodes[i].Attributes)
		jAttr := types.Attributes(nodes[j].Attributes)

		iMount, iExist := iAttr.Find(storageAttributeMount).AsString()
		jMount, jExist := jAttr.Find(storageAttributeMount).AsString()

		if !iExist {
			return true
		} else if !jExist {
			return false
		}

		return iMount < jMount
	})

	*sdl = nodes

	return nil
}

func (sdl *v2StorageAttributes) UnmarshalYAML(node *yaml.Node) error {
	var attr v2StorageAttributes

	var res map[string]string

	if err := node.Decode(&res); err != nil {
		return err
	}

	for k, v := range res {
		if _, validKey := allowedStorageAttributes[k]; !validKey {
			return errors.Wrap(errUnsupportedStorageAttribute, k)
		}

		attr = append(attr, types.Attribute{
			Key:   k,
			Value: v,
		})
	}

	// at this point keys are unique in attributes parsed from sdl so don't need to use sort.SliceStable
	sort.Slice(attr, func(i, j int) bool {
		return attr[i].Key < attr[j].Key
	})

	*sdl = attr

	return nil
}
