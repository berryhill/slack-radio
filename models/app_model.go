package models

import ()

type App struct {
  Active                    bool
  ActiveListenerId          []*bson.ObjectId
  activeListeners           []*Listener
  ActiveSpotifyAccountIds   []*SpotifyAccount
  activeSpotifyAccounts     []*bson.ObjectId
  Id                        *bson.ObjectId
  ListenerIds               []*bson.ObjectId
  listeners                 []*Listener
  Player                    *Player
  SpofityAccountIds         []*bson.ObjectIds
  spotifyAccounts           []*SpotifyAccount
}

type Listener stuct {
  Id                        *bson.ObjectId
  SpotifyAccountId          *bson.ObjectIds
  spotifyAccount            *SpotifyAccount
  Username                  string
}

type Play struct {
  Id                        *bson.ObjectId
  Timestamp                 time.Timestamp
}

type Player struct {
  CurrentTrack              *Track
  QueuedTracks              []*Track
  State                     string
}

func (p *Player) AddToQueue(track *Track) error {

  // TODO: implement

  return nil
}

func (p *Player) Pause() error {

  // TODO: implement

  return nil
}

func (p *Player) Play() error {

  // TODO: implement

  return nil
}

func (p *Player) SeekForward() error {

  // TODO: implement

  return nil
}

type SpotifyAccount struct {
  Id                        *bson.ObjectId
  SpotifyId                 string
  Username                  string
}

type Track struct {
  Artist                    string
  DownVotes                 int
  Genre                     string
  Id                        *bson.ObjectId
  Name                      string
  Plays                     []*Play
  SpotifyId                 string
  UpVotes                   int
}
