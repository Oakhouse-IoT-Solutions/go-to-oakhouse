# Blog API Example

This example demonstrates how to build a complete blog API using the Go To Oakhouse framework. The API includes user authentication, blog posts, comments, and categories.

## Features

- **User Management**: Registration, authentication, and profile management
- **Blog Posts**: Create, read, update, and delete blog posts
- **Comments**: Add comments to blog posts
- **Categories**: Organize posts by categories
- **Authentication**: JWT-based authentication
- **Authorization**: Role-based access control
- **Pagination**: Paginated API responses
- **Filtering**: Filter posts by category, author, status
- **Search**: Full-text search in posts
- **File Upload**: Upload images for posts

## Quick Start

### 1. Create the Project

```bash
# Create new project
oakhouse new blog-api
cd blog-api
```

### 2. Generate Resources

```bash
# Generate User resource
oakhouse generate resource User name:string email:string password:string role:string avatar_url:string bio:text

# Generate Category resource
oakhouse generate resource Category name:string slug:string description:text

# Generate Post resource
oakhouse generate resource Post title:string slug:string content:text excerpt:string author_id:uuid category_id:uuid featured_image:string status:string published_at:time

# Generate Comment resource
oakhouse generate resource Comment content:text post_id:uuid author_id:uuid parent_id:uuid status:string

# Generate Tag resource
oakhouse generate resource Tag name:string slug:string

# Generate PostTag junction table
oakhouse generate model PostTag post_id:uuid tag_id:uuid
```

### 3. Set Up Database

```bash
# Configure your .env file
cp .env.example .env
# Edit .env with your database credentials

# Run migrations
oakhouse migrate up
```

### 4. Start Development Server

```bash
oakhouse serve
```

## API Endpoints

### Authentication

```
POST   /api/v1/auth/register     # Register new user
POST   /api/v1/auth/login        # Login user
POST   /api/v1/auth/logout       # Logout user
POST   /api/v1/auth/refresh      # Refresh JWT token
GET    /api/v1/auth/me           # Get current user profile
PUT    /api/v1/auth/profile      # Update user profile
POST   /api/v1/auth/change-password  # Change password
```

### Users

```
GET    /api/v1/users             # List users (admin only)
GET    /api/v1/users/:id         # Get user by ID
PUT    /api/v1/users/:id         # Update user (admin only)
DELETE /api/v1/users/:id         # Delete user (admin only)
GET    /api/v1/users/:id/posts   # Get user's posts
```

### Categories

```
GET    /api/v1/categories        # List categories
GET    /api/v1/categories/:id    # Get category by ID
POST   /api/v1/categories        # Create category (admin only)
PUT    /api/v1/categories/:id    # Update category (admin only)
DELETE /api/v1/categories/:id    # Delete category (admin only)
GET    /api/v1/categories/:id/posts  # Get category posts
```

### Posts

```
GET    /api/v1/posts             # List posts
GET    /api/v1/posts/:id         # Get post by ID
POST   /api/v1/posts             # Create post (authenticated)
PUT    /api/v1/posts/:id         # Update post (author/admin)
DELETE /api/v1/posts/:id         # Delete post (author/admin)
GET    /api/v1/posts/:id/comments # Get post comments
POST   /api/v1/posts/:id/publish  # Publish post (author/admin)
POST   /api/v1/posts/:id/unpublish # Unpublish post (author/admin)
```

### Comments

```
GET    /api/v1/comments          # List comments (admin only)
GET    /api/v1/comments/:id      # Get comment by ID
POST   /api/v1/comments          # Create comment (authenticated)
PUT    /api/v1/comments/:id      # Update comment (author/admin)
DELETE /api/v1/comments/:id      # Delete comment (author/admin)
POST   /api/v1/comments/:id/approve   # Approve comment (admin)
POST   /api/v1/comments/:id/reject    # Reject comment (admin)
```

### Tags

```
GET    /api/v1/tags              # List tags
GET    /api/v1/tags/:id          # Get tag by ID
POST   /api/v1/tags              # Create tag (admin only)
PUT    /api/v1/tags/:id          # Update tag (admin only)
DELETE /api/v1/tags/:id          # Delete tag (admin only)
GET    /api/v1/tags/:id/posts    # Get tag posts
```

### Search

```
GET    /api/v1/search/posts      # Search posts
GET    /api/v1/search/users      # Search users
```

### File Upload

```
POST   /api/v1/upload/image      # Upload image
POST   /api/v1/upload/avatar     # Upload user avatar
```

## Sample Requests

### Register a New User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "bio": "A passionate blogger"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create a Category

```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Technology",
    "slug": "technology",
    "description": "Posts about technology and programming"
  }'
```

### Create a Blog Post

```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Getting Started with Go",
    "slug": "getting-started-with-go",
    "content": "Go is a programming language developed by Google...",
    "excerpt": "Learn the basics of Go programming language",
    "category_id": "CATEGORY_UUID",
    "status": "draft",
    "tags": ["go", "programming", "tutorial"]
  }'
```

### Get Posts with Filtering

```bash
# Get published posts with pagination
curl "http://localhost:8080/api/v1/posts?status=published&page=1&page_size=10&sort_by=published_at&sort_order=desc"

# Get posts by category
curl "http://localhost:8080/api/v1/posts?category_id=CATEGORY_UUID"

# Get posts by author
curl "http://localhost:8080/api/v1/posts?author_id=USER_UUID"

# Search posts
curl "http://localhost:8080/api/v1/search/posts?q=golang&page=1&page_size=10"
```

### Add a Comment

```bash
curl -X POST http://localhost:8080/api/v1/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "content": "Great article! Thanks for sharing.",
    "post_id": "POST_UUID"
  }'
```

### Reply to a Comment

```bash
curl -X POST http://localhost:8080/api/v1/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "content": "Thank you for the feedback!",
    "post_id": "POST_UUID",
    "parent_id": "PARENT_COMMENT_UUID"
  }'
```

## Data Models

### User Model

```go
type User struct {
    ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Name      string         `gorm:"not null" json:"name"`
    Email     string         `gorm:"uniqueIndex;not null" json:"email"`
    Password  string         `gorm:"not null" json:"-"`
    Role      string         `gorm:"default:user" json:"role"` // user, admin, moderator
    AvatarURL string         `json:"avatar_url"`
    Bio       string         `gorm:"type:text" json:"bio"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    
    // Relationships
    Posts    []Post    `gorm:"foreignKey:AuthorID" json:"posts,omitempty"`
    Comments []Comment `gorm:"foreignKey:AuthorID" json:"comments,omitempty"`
}
```

### Post Model

```go
type Post struct {
    ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Title         string         `gorm:"not null" json:"title"`
    Slug          string         `gorm:"uniqueIndex;not null" json:"slug"`
    Content       string         `gorm:"type:text" json:"content"`
    Excerpt       string         `gorm:"type:text" json:"excerpt"`
    AuthorID      uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`
    CategoryID    *uuid.UUID     `gorm:"type:uuid" json:"category_id"`
    FeaturedImage string         `json:"featured_image"`
    Status        string         `gorm:"default:draft" json:"status"` // draft, published, archived
    ViewCount     int            `gorm:"default:0" json:"view_count"`
    PublishedAt   *time.Time     `json:"published_at"`
    CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    
    // Relationships
    Author   User      `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
    Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
    Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
    Tags     []Tag     `gorm:"many2many:post_tags" json:"tags,omitempty"`
}
```

### Comment Model

```go
type Comment struct {
    ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Content   string         `gorm:"type:text;not null" json:"content"`
    PostID    uuid.UUID      `gorm:"type:uuid;not null" json:"post_id"`
    AuthorID  uuid.UUID      `gorm:"type:uuid;not null" json:"author_id"`
    ParentID  *uuid.UUID     `gorm:"type:uuid" json:"parent_id"`
    Status    string         `gorm:"default:pending" json:"status"` // pending, approved, rejected
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    
    // Relationships
    Post     Post      `gorm:"foreignKey:PostID" json:"post,omitempty"`
    Author   User      `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
    Parent   *Comment  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
    Replies  []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
}
```

### Category Model

```go
type Category struct {
    ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Name        string         `gorm:"not null" json:"name"`
    Slug        string         `gorm:"uniqueIndex;not null" json:"slug"`
    Description string         `gorm:"type:text" json:"description"`
    PostCount   int            `gorm:"default:0" json:"post_count"`
    CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    
    // Relationships
    Posts []Post `gorm:"foreignKey:CategoryID" json:"posts,omitempty"`
}
```

### Tag Model

```go
type Tag struct {
    ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Name      string         `gorm:"not null" json:"name"`
    Slug      string         `gorm:"uniqueIndex;not null" json:"slug"`
    PostCount int            `gorm:"default:0" json:"post_count"`
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    
    // Relationships
    Posts []Post `gorm:"many2many:post_tags" json:"posts,omitempty"`
}
```

## Advanced Features

### Custom Scopes

The example includes custom query scopes for complex filtering:

```go
// scope/post/filter_by_status.go
func FilterByStatus(status string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if status == "" {
            return db
        }
        return db.Where("status = ?", status)
    }
}

// scope/post/filter_by_published.go
func FilterByPublished() func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("status = ? AND published_at IS NOT NULL AND published_at <= ?", "published", time.Now())
    }
}

// scope/post/search_content.go
func SearchContent(query string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if query == "" {
            return db
        }
        searchTerm := "%" + query + "%"
        return db.Where("title ILIKE ? OR content ILIKE ? OR excerpt ILIKE ?", searchTerm, searchTerm, searchTerm)
    }
}
```

### Custom Middleware

```go
// middleware/post_ownership.go
func PostOwnership() fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("user_id").(uuid.UUID)
        userRole := c.Locals("user_role").(string)
        postID := c.Params("id")
        
        // Admin can access any post
        if userRole == "admin" {
            return c.Next()
        }
        
        // Check if user owns the post
        var post entity.Post
        if err := db.First(&post, "id = ? AND author_id = ?", postID, userID).Error; err != nil {
            return util.SendError(c, fiber.StatusForbidden, "Access denied", nil)
        }
        
        return c.Next()
    }
}
```

### File Upload Service

```go
// service/upload_service.go
type UploadService interface {
    UploadImage(file *multipart.FileHeader) (string, error)
    DeleteImage(url string) error
}

type uploadService struct {
    config *config.Config
}

func (s *uploadService) UploadImage(file *multipart.FileHeader) (string, error) {
    // Validate file type
    if !isValidImageType(file.Header.Get("Content-Type")) {
        return "", errors.New("invalid file type")
    }
    
    // Generate unique filename
    filename := generateUniqueFilename(file.Filename)
    
    // Save file to storage (local, S3, etc.)
    if err := saveFile(file, filename); err != nil {
        return "", err
    }
    
    return s.config.BaseURL + "/uploads/" + filename, nil
}
```

## Testing

The example includes comprehensive tests:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test -tags=integration ./...
```

## Deployment

### Docker

```bash
# Build and run with Docker Compose
docker-compose up -d

# View logs
docker-compose logs -f app
```

### Production Environment

```bash
# Build for production
oakhouse build

# Run migrations
APP_ENV=production ./bin/blog-api migrate up

# Start production server
APP_ENV=production ./bin/blog-api
```

This example demonstrates the power and simplicity of the Go To Oakhouse framework for building modern web APIs. The generated code follows best practices and provides a solid foundation for building scalable applications.
