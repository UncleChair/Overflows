package consts

const (
	AppEnv             = "prod"                                // Server stage
	Version            = "v0.0.1"                              // Server version
	IVKey              = "rpFObp7BcJ2t7wgp"                    // AES key
	JWTKey             = "9z6t1iRvyC8DGbFB"                    // JWT key
	PublicPath         = "./"                                  // Public storage path
	AvatarFolder       = "avatars"                             // Avatar storage folder
	AvatarPath         = PublicPath + "/" + AvatarFolder       // Avatar storage path
	ProjectCoverFolder = "project_covers"                      // Project cover storage folder
	ProjectCoverPath   = PublicPath + "/" + ProjectCoverFolder // Project cover storage path
)
