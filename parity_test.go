package myrtle_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/gzuidhof/myrtle"
)

// builderOnlyAddMethods lists Builder Add* methods that are intentionally not
// mirrored on Group. Keep this empty unless there is a deliberate reason a
// method cannot exist on a Group; every entry should be justified by a comment.
var builderOnlyAddMethods = map[string]string{}

// groupOnlyAddMethods lists Group Add* methods that are intentionally not
// mirrored on Builder. Keep this empty unless there is a deliberate reason a
// method cannot exist on a Builder; every entry should be justified by a comment.
var groupOnlyAddMethods = map[string]string{}

// TestGroupBuilderAddParity enforces that *Builder and *Group expose the same
// set of chaining Add* methods with identical parameter lists (the only
// permitted difference is the receiver and the *Builder vs *Group return type).
//
// Parity is checked in both directions: adding an AddX method to one type
// without adding it to the other fails the test. This guards the documented
// invariant that a Group supports the same block APIs as top-level builder
// chaining. When you add an AddX method to one type, add the matching method to
// the other (or, with justification, list it in the corresponding allowlist).
func TestGroupBuilderAddParity(t *testing.T) {
	builderType := reflect.TypeOf(&myrtle.Builder{})
	groupType := reflect.TypeOf(&myrtle.Group{})

	assertAddParity(t, builderType, groupType, "Builder", "Group", "builder_group.go", builderOnlyAddMethods)
	assertAddParity(t, groupType, builderType, "Group", "Builder", "builder_blocks.go", groupOnlyAddMethods)
}

// assertAddParity checks that every Add* method on srcType has a matching method
// on dstType with an identical (non-receiver) parameter list, skipping any names
// in allowlist. srcName/dstName are used for messages; addFile names the file
// where the missing method should typically be added.
func assertAddParity(t *testing.T, srcType, dstType reflect.Type, srcName, dstName, addFile string, allowlist map[string]string) {
	t.Helper()

	for i := 0; i < srcType.NumMethod(); i++ {
		method := srcType.Method(i)
		if !strings.HasPrefix(method.Name, "Add") {
			continue
		}
		if reason, ok := allowlist[method.Name]; ok {
			t.Logf("skipping %s-only method %s (%s)", srcName, method.Name, reason)
			continue
		}

		dstMethod, ok := dstType.MethodByName(method.Name)
		if !ok {
			t.Errorf("%s is missing method %s present on %s; add it to %s or list it in the corresponding allowlist", dstName, method.Name, srcName, addFile)
			continue
		}

		if diff := paramSignatureDiff(method.Func.Type(), dstMethod.Func.Type(), srcName, dstName); diff != "" {
			t.Errorf("%s.%s parameters do not match %s.%s: %s", dstName, method.Name, srcName, method.Name, diff)
		}
	}
}

// paramSignatureDiff compares the non-receiver parameters of two methods and
// returns a human-readable description of the first mismatch, or "" if the
// parameter lists are identical. The receiver (parameter 0) is ignored so that
// *Builder and *Group receivers are treated as equivalent.
func paramSignatureDiff(srcFn, dstFn reflect.Type, srcName, dstName string) string {
	if srcFn.IsVariadic() != dstFn.IsVariadic() {
		return "one method is variadic and the other is not"
	}

	srcParams := methodParams(srcFn)
	dstParams := methodParams(dstFn)

	if len(srcParams) != len(dstParams) {
		return srcName + " has " + joinTypes(srcParams) + "; " + dstName + " has " + joinTypes(dstParams)
	}

	for i := range srcParams {
		if srcParams[i] != dstParams[i] {
			return "param " + strconv.Itoa(i) + ": " + srcName + "=" + srcParams[i] + " " + dstName + "=" + dstParams[i]
		}
	}

	return ""
}

func methodParams(fn reflect.Type) []string {
	params := make([]string, 0, fn.NumIn())
	for i := 1; i < fn.NumIn(); i++ { // skip receiver at index 0
		params = append(params, fn.In(i).String())
	}
	return params
}

func joinTypes(types []string) string {
	if len(types) == 0 {
		return "()"
	}
	return "(" + strings.Join(types, ", ") + ")"
}
