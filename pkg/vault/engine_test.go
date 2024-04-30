package vault

import (
	"github.com/stretchr/testify/require"
)

func (s *VaultSuite) TestEnableKV2EngineErrorIfNotForced() {
	testCases := []struct {
		name    string
		force   bool
		path    string
		prepare bool
		err     bool
	}{
		{
			name:  "engine does not exist, no force",
			force: false,
			path:  "case-1",
			err:   false,
		},
		{
			name:    "engine does exist, no force",
			force:   false,
			prepare: true,
			path:    "case-2",
			err:     true,
		},
		{
			name:    "engine does exist, force",
			force:   true,
			prepare: true,
			path:    "case-3",
			err:     false,
		},
		{
			name:    "engine does exist, no force",
			force:   false,
			prepare: true,
			path:    "case-4",
			err:     true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if tc.prepare {
				s.Require().NoError(s.client.EnableKV2Engine(tc.path))
			}

			err := s.client.EnableKV2EngineErrorIfNotForced(tc.force, tc.path)

			s.Require().Equal(tc.err, err != nil, tc.name)
		})
	}
}

func (s *VaultSuite) TestListAllKVSecretEngines() {
	testCases := []struct {
		name     string
		engines  []string
		expected Engines
	}{
		{
			name:    "test",
			engines: []string{"1", "2", "3"},
			expected: Engines{
				"": []string{"secret/", "1/", "2/", "3/"}, // enabled by default
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			for _, engine := range tc.engines {
				require.NoError(s.Suite.T(), s.client.EnableKV2Engine(engine), tc.name)
			}

			res, err := s.client.ListAllKVSecretEngines("")
			s.Require().NoError(err, tc.name)

			s.Require().ElementsMatch(tc.expected[""], res[""], tc.name)
		})
	}
}
