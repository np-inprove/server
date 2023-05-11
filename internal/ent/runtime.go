// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/np-inprove/server/internal/ent/academicschool"
	"github.com/np-inprove/server/internal/ent/accessory"
	"github.com/np-inprove/server/internal/ent/course"
	"github.com/np-inprove/server/internal/ent/deadline"
	"github.com/np-inprove/server/internal/ent/event"
	"github.com/np-inprove/server/internal/ent/forumpost"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/milestone"
	"github.com/np-inprove/server/internal/ent/reaction"
	"github.com/np-inprove/server/internal/ent/schema"
	"github.com/np-inprove/server/internal/ent/studyplan"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/ent/userpet"
	"github.com/np-inprove/server/internal/ent/voucher"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	academicschoolFields := schema.AcademicSchool{}.Fields()
	_ = academicschoolFields
	// academicschoolDescName is the schema descriptor for name field.
	academicschoolDescName := academicschoolFields[0].Descriptor()
	// academicschool.NameValidator is a validator for the "name" field. It is called by the builders before save.
	academicschool.NameValidator = academicschoolDescName.Validators[0].(func(string) error)
	// academicschoolDescSchoolCode is the schema descriptor for school_code field.
	academicschoolDescSchoolCode := academicschoolFields[1].Descriptor()
	// academicschool.SchoolCodeValidator is a validator for the "school_code" field. It is called by the builders before save.
	academicschool.SchoolCodeValidator = academicschoolDescSchoolCode.Validators[0].(func(string) error)
	accessoryMixin := schema.Accessory{}.Mixin()
	accessoryMixinFields0 := accessoryMixin[0].Fields()
	_ = accessoryMixinFields0
	accessoryFields := schema.Accessory{}.Fields()
	_ = accessoryFields
	// accessoryDescName is the schema descriptor for name field.
	accessoryDescName := accessoryMixinFields0[0].Descriptor()
	// accessory.NameValidator is a validator for the "name" field. It is called by the builders before save.
	accessory.NameValidator = accessoryDescName.Validators[0].(func(string) error)
	// accessoryDescPointsRequired is the schema descriptor for points_required field.
	accessoryDescPointsRequired := accessoryMixinFields0[2].Descriptor()
	// accessory.PointsRequiredValidator is a validator for the "points_required" field. It is called by the builders before save.
	accessory.PointsRequiredValidator = accessoryDescPointsRequired.Validators[0].(func(int) error)
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescName is the schema descriptor for name field.
	courseDescName := courseFields[0].Descriptor()
	// course.NameValidator is a validator for the "name" field. It is called by the builders before save.
	course.NameValidator = courseDescName.Validators[0].(func(string) error)
	// courseDescCourseID is the schema descriptor for course_id field.
	courseDescCourseID := courseFields[1].Descriptor()
	// course.CourseIDValidator is a validator for the "course_id" field. It is called by the builders before save.
	course.CourseIDValidator = courseDescCourseID.Validators[0].(func(string) error)
	deadlineFields := schema.Deadline{}.Fields()
	_ = deadlineFields
	// deadlineDescName is the schema descriptor for name field.
	deadlineDescName := deadlineFields[0].Descriptor()
	// deadline.NameValidator is a validator for the "name" field. It is called by the builders before save.
	deadline.NameValidator = deadlineDescName.Validators[0].(func(string) error)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescName is the schema descriptor for name field.
	eventDescName := eventFields[0].Descriptor()
	// event.NameValidator is a validator for the "name" field. It is called by the builders before save.
	event.NameValidator = eventDescName.Validators[0].(func(string) error)
	forumpostFields := schema.ForumPost{}.Fields()
	_ = forumpostFields
	// forumpostDescTitle is the schema descriptor for title field.
	forumpostDescTitle := forumpostFields[0].Descriptor()
	// forumpost.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	forumpost.TitleValidator = forumpostDescTitle.Validators[0].(func(string) error)
	// forumpostDescContent is the schema descriptor for content field.
	forumpostDescContent := forumpostFields[1].Descriptor()
	// forumpost.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	forumpost.ContentValidator = forumpostDescContent.Validators[0].(func(string) error)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescPath is the schema descriptor for path field.
	groupDescPath := groupFields[0].Descriptor()
	// group.PathValidator is a validator for the "path" field. It is called by the builders before save.
	group.PathValidator = groupDescPath.Validators[0].(func(string) error)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[1].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	institutionFields := schema.Institution{}.Fields()
	_ = institutionFields
	// institutionDescName is the schema descriptor for name field.
	institutionDescName := institutionFields[0].Descriptor()
	// institution.NameValidator is a validator for the "name" field. It is called by the builders before save.
	institution.NameValidator = institutionDescName.Validators[0].(func(string) error)
	milestoneFields := schema.Milestone{}.Fields()
	_ = milestoneFields
	// milestoneDescName is the schema descriptor for name field.
	milestoneDescName := milestoneFields[0].Descriptor()
	// milestone.NameValidator is a validator for the "name" field. It is called by the builders before save.
	milestone.NameValidator = milestoneDescName.Validators[0].(func(string) error)
	reactionFields := schema.Reaction{}.Fields()
	_ = reactionFields
	// reactionDescEmoji is the schema descriptor for emoji field.
	reactionDescEmoji := reactionFields[2].Descriptor()
	// reaction.EmojiValidator is a validator for the "emoji" field. It is called by the builders before save.
	reaction.EmojiValidator = reactionDescEmoji.Validators[0].(func(string) error)
	studyplanFields := schema.StudyPlan{}.Fields()
	_ = studyplanFields
	// studyplanDescName is the schema descriptor for name field.
	studyplanDescName := studyplanFields[0].Descriptor()
	// studyplan.NameValidator is a validator for the "name" field. It is called by the builders before save.
	studyplan.NameValidator = studyplanDescName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[0].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[1].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPasswordHash is the schema descriptor for password_hash field.
	userDescPasswordHash := userFields[3].Descriptor()
	// user.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	user.PasswordHashValidator = userDescPasswordHash.Validators[0].(func(string) error)
	// userDescPoints is the schema descriptor for points field.
	userDescPoints := userFields[4].Descriptor()
	// user.PointsValidator is a validator for the "points" field. It is called by the builders before save.
	user.PointsValidator = userDescPoints.Validators[0].(func(int) error)
	// userDescPointsAwardedCount is the schema descriptor for points_awarded_count field.
	userDescPointsAwardedCount := userFields[5].Descriptor()
	// user.PointsAwardedCountValidator is a validator for the "points_awarded_count" field. It is called by the builders before save.
	user.PointsAwardedCountValidator = userDescPointsAwardedCount.Validators[0].(func(int) error)
	userpetFields := schema.UserPet{}.Fields()
	_ = userpetFields
	// userpetDescHungerPercentage is the schema descriptor for hunger_percentage field.
	userpetDescHungerPercentage := userpetFields[0].Descriptor()
	// userpet.HungerPercentageValidator is a validator for the "hunger_percentage" field. It is called by the builders before save.
	userpet.HungerPercentageValidator = func() func(float64) error {
		validators := userpetDescHungerPercentage.Validators
		fns := [...]func(float64) error{
			validators[0].(func(float64) error),
			validators[1].(func(float64) error),
		}
		return func(hunger_percentage float64) error {
			for _, fn := range fns {
				if err := fn(hunger_percentage); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userpetDescEnabledSvgGroupElementIds is the schema descriptor for enabled_svg_group_element_ids field.
	userpetDescEnabledSvgGroupElementIds := userpetFields[1].Descriptor()
	// userpet.DefaultEnabledSvgGroupElementIds holds the default value on creation for the enabled_svg_group_element_ids field.
	userpet.DefaultEnabledSvgGroupElementIds = userpetDescEnabledSvgGroupElementIds.Default.(map[string]bool)
	voucherMixin := schema.Voucher{}.Mixin()
	voucherMixinFields0 := voucherMixin[0].Fields()
	_ = voucherMixinFields0
	voucherFields := schema.Voucher{}.Fields()
	_ = voucherFields
	// voucherDescName is the schema descriptor for name field.
	voucherDescName := voucherMixinFields0[0].Descriptor()
	// voucher.NameValidator is a validator for the "name" field. It is called by the builders before save.
	voucher.NameValidator = voucherDescName.Validators[0].(func(string) error)
	// voucherDescPointsRequired is the schema descriptor for points_required field.
	voucherDescPointsRequired := voucherMixinFields0[2].Descriptor()
	// voucher.PointsRequiredValidator is a validator for the "points_required" field. It is called by the builders before save.
	voucher.PointsRequiredValidator = voucherDescPointsRequired.Validators[0].(func(int) error)
}
