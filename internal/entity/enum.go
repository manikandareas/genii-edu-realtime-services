package entity

type RoleEnum string
type ClassAccessTypeEnum string
type ClassCompletionStatusEnum string
type StudentProgressTypeEnum string
type FilesTypeEnum string

const (
	TeacherRole          RoleEnum                  = "teacher"
	StudentRole          RoleEnum                  = "student"
	PublicAccess         ClassAccessTypeEnum       = "public"
	PrivateAccess        ClassAccessTypeEnum       = "private"
	OngoingStatus        ClassCompletionStatusEnum = "ongoing"
	CompletedStatus      ClassCompletionStatusEnum = "completed"
	ArchivedStatus       ClassCompletionStatusEnum = "archived"
	AssignmentCompletion StudentProgressTypeEnum   = "assignment_completion"
	MaterialView         StudentProgressTypeEnum   = "material_view"
	ImageFile            FilesTypeEnum             = "image"
	VideoFile            FilesTypeEnum             = "video"
	PDFFile              FilesTypeEnum             = "pdf"
	YouTubeFile          FilesTypeEnum             = "youtube"
)
