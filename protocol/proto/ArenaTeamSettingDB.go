// Code generated by gucooing DO NOT EDIT.

package proto

type ArenaTeamSettingDB struct{
    EchelonType EchelonType
    LeaderCharacterId int64
    TSSInteractionCharacterId int64
    MainCharacters []*ArenaCharacterDB
    SupportCharacters []*ArenaCharacterDB
    TSSCharacterDB *ArenaCharacterDB
    MapId int64
}

