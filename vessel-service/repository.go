// vessel-service/respository.go

package main

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection.Insert(vessel)
}
