package dockercompose_test

import (
	"testing"

	"github.com/Bocmah/phpdocker-gen/internal/dockercompose"
)

func TestRestartPolicy_Render(t *testing.T) {
	tests := map[string]struct {
		input dockercompose.RestartPolicy
		want  string
	}{
		"policy always": {
			input: dockercompose.RestartPolicyAlways,
			want:  "restart: always",
		},
		"policy no": {
			input: dockercompose.RestartPolicyNo,
			want:  `restart: "no"`,
		},
		"unknown policy": {
			input: dockercompose.RestartPolicy("some-policy"),
			want:  "",
		},
		"empty policy": {
			input: dockercompose.RestartPolicy(""),
			want:  "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.Render()
			if tc.want != got {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
