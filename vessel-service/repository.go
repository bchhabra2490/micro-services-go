// vessel-service/respository.go

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection.Insert(vessel)
}