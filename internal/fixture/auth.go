package fixture

import (
	"encoding/json"
	"github.com/np-inprove/server/internal/entity"
	"github.com/np-inprove/server/internal/entity/institution"
	"github.com/np-inprove/server/internal/hash"
	"time"
)

var (
	encodedPasswordString = "{\"h\":\"15BrHhXrHhYM2yMFMrghE9gnY950dHNF1klveR5EdeY=\",\"s\":\"J7lr7qPnc/S4iGgdJQWSK2xVELHLE9N4+KIdcdBdHhM=\",\"t\":1,\"m\":65536,\"k\":32}"
)

var (
	AuthJWKString = "{\"kty\": \"EC\",\"d\": \"AFDDClQjKNoNfFGZB9iLP2gRmAJXBa-mp_XIdYMJjNKLIRELS03WArued737uRt9ayopD6AEQZOiRSLs8o9LrT5B\",\"use\": \"sig\",\"crv\": \"P-521\",\"x\": \"AZGgvUEwezY2RHU7l4m1aF5vp8Vj8CUo_nyqUzVVwoVjFz_mbZM4ZznXrAvuLKlAqagI2bHHj-97W7Blp3VaI0A5\",\"y\": \"AdmmlsSae6qoGTPsUYj7gSOGMJcaRdHpcyvQ3hdq5owJnyozbRib5mUYXwsV_voIKBTSkeztQSRagiILwqiEy5-y\",\"alg\": \"ES512\"}"

	InstitutionInviteLinkNP = &entity.InstitutionInviteLink{
		ID:   1,
		Code: "abcdef",
		Role: institution.RoleMember,
		Edges: entity.InstitutionInviteLinkEdges{
			Institution: InstitutionNP,
		},
	}

	Institutions = []*entity.Institution{InstitutionNP, InstitutionSP}

	InstitutionNP = &entity.Institution{
		ID:          1,
		Name:        "Ngee Ann Polytechnic",
		ShortName:   "NP",
		Description: "Ngee Ann Polytechnic",
	}
	InstitutionSP = &entity.Institution{
		ID:          1,
		Name:        "Singapore Polytechnic",
		ShortName:   "SP",
		Description: "Singapore Polytechnic",
	}

	UserJohnPassword = "example"
	UserJohn         = &entity.User{
		ID:        1,
		FirstName: "John",
		LastName:  "The Ripper",
		Email:     "john@theripper.com",
		Password: (func() (enc hash.Encoded) {
			_ = json.Unmarshal([]byte(encodedPasswordString), &enc)
			return
		})(),
		Points:                 1,
		PointsAwardedCount:     1,
		PointsAwardedResetTime: time.Now(),
		GodMode:                false,
		Role:                   institution.RoleMember,
		Edges: entity.UserEdges{
			Institution: InstitutionNP,
		},
	}
)
