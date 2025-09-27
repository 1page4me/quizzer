# Overview

QuizGen is now a cross-platform mobile game built with Flutter, featuring AI-powered educational quizzes with gamification elements. The system uses a Go backend with Firebase integration to provide real-time multiplayer capabilities, cloud storage, and comprehensive analytics. Students can take timed quizzes with difficulty-based countdown timers, earn points and achievements, compete with friends, and track their learning progress across Android, iOS, Web, and Windows platforms.

# User Preferences

Preferred communication style: Simple, everyday language.
Target users: Young children (ages 3-5) with parent/teacher-controlled content management.
UX Principles: Child-friendly interface, simplified workflows, gamification elements.
Code Architecture: DRY (Don't Repeat Yourself), Singleton patterns, DOT (Do One Thing) functions.
Authentication: Role-based access control with cute game-style login (student/teacher/admin roles).

# System Architecture

## Full-Stack TypeScript Architecture
The application uses a modern full-stack TypeScript architecture with strict type safety across client, server, and shared components. ESM modules are used throughout for better performance and modern JavaScript standards.

## Mobile Frontend (Flutter)
- **Framework**: Flutter 3.0+ for cross-platform development (Android, iOS, Web, Windows)
- **State Management**: Riverpod for reactive state management
- **UI Components**: Material Design 3 with custom gaming themes
- **Animations**: Lottie animations and custom Flutter animations
- **Local Storage**: Hive for offline data persistence
- **HTTP Client**: Dio for API communication with interceptors
- **Gaming Features**: Haptic feedback, sound effects, particle animations

## Backend (Go)
- **Framework**: Gin for high-performance HTTP routing
- **Concurrency**: Go routines for multithreading and real-time features
- **Database**: Firebase Firestore for NoSQL document storage
- **Real-time**: Firebase Realtime Database for live multiplayer features
- **Authentication**: Firebase Auth with social login support
- **File Storage**: Firebase Cloud Storage for syllabus materials
- **Functions**: Firebase Cloud Functions for serverless operations

## Database Layer (Firebase)
- **Firestore**: NoSQL document database for user data, quizzes, and analytics
- **Realtime Database**: Real-time synchronization for multiplayer features
- **Cloud Storage**: File storage for admin-uploaded content and user assets
- **Security Rules**: Firebase security rules for data protection
- **Indexing**: Composite indexes for efficient queries
- **Offline Support**: Local caching and sync when connection restored

## Data Models (Updated Architecture)
- **Users**: Role-based system (admin/teacher/student) with authentication
- **Students**: Extended profiles with learning styles and year group tracking
- **Subjects**: Admin-controlled curriculum subjects with year group mapping
- **Topics**: Hierarchical content organization under subjects
- **Syllabus Content**: Admin-uploaded materials linked to specific topics
- **Quizzes**: AI-generated from admin-curated content with enhanced metadata
- **Quiz Attempts**: Performance tracking with streak counting and percentile ranking
- **Achievements**: Gamified progress system with multiple achievement types

## AI Integration
- **OpenAI GPT-4o**: Latest model for quiz question generation
- **Content Analysis**: Automatic topic extraction from uploaded files
- **Adaptive Difficulty**: Questions generated based on student level and performance
- **Natural Language Processing**: Content parsing and question formulation

## Content Management System
- **Admin-Only Upload**: Teachers and admins control all curriculum content
- **Structured Content**: Topics organized by subject, year group, and difficulty
- **Student Selection**: Students choose from pre-approved topics only
- **Quality Control**: All content vetted before student access

## Authentication & Security
- **Firebase Auth**: Role-based authentication with social login support
- **Role-Based Access Control (RBAC)**: Student, teacher, and admin permissions
- **Route Protection**: Navigation guards based on user roles
- **Game-Style Login**: Cute, engaging interface resembling mobile games
- **Security Rules**: Firebase security rules for data protection
- **API Security**: Input validation and sanitization

## Performance Features
- **Query Optimization**: Efficient database queries with proper indexing
- **Caching**: Query caching with TanStack Query
- **Bundle Optimization**: Code splitting and tree shaking with Vite
- **Asset Management**: Optimized static asset serving

## Development Environment
- **Hot Reload**: Vite HMR for instant development feedback
- **Type Checking**: Strict TypeScript configuration
- **Path Aliases**: Clean import paths with barrel exports
- **Error Overlay**: Runtime error modal for development debugging

# External Dependencies

## Firebase Services
- **Firebase Auth**: User authentication with social providers (Google, Apple, Facebook)
- **Firestore**: Document database for scalable data storage
- **Realtime Database**: Real-time synchronization for live features
- **Cloud Storage**: File storage with CDN distribution
- **Cloud Functions**: Serverless functions for quiz generation and processing
- **Analytics**: User behavior tracking and performance analytics
- **Crashlytics**: Real-time crash reporting and debugging

## AI Services  
- **OpenAI API**: GPT-4o model integration via Firebase Cloud Functions
- **Firebase ML**: On-device ML models for offline features

## Component Architecture (DRY/DOT Principles)
- **Atomic Components**: Single-purpose components (LoadingSpinner, DifficultyBadge, ScoreDisplay)
- **Molecular Components**: Combined atoms (SubjectCard, TopicSelector, StudentQuizConfig)
- **Singleton Services**: UserRoleManager for permission management
- **Constants Library**: Centralized configuration values and UI constants
- **Radix UI**: Accessible component primitives
- **Shadcn/ui**: Consistent design system
- **Recharts**: Analytics visualization

## Development Tools
- **Replit Integration**: Development environment with live preview and collaborative features
- **Vite Plugins**: Runtime error overlay and cartographer for enhanced debugging

## File Processing
- **Multer**: Express middleware for handling multipart/form-data uploads
- **File Type Validation**: MIME type checking and file size limits

## Styling & Fonts
- **Tailwind CSS**: Utility-first CSS framework with custom design tokens
- **Google Fonts**: Typography system with multiple font families (DM Sans, Fira Code, Geist Mono)

## Architecture Principles Applied
- **DRY (Don't Repeat Yourself)**: Centralized constants, reusable atomic components
- **Singleton Pattern**: UserRoleManager for global state management
- **DOT (Do One Thing)**: Each component has single responsibility
- **Student-Centric UX**: Simplified interfaces, guided workflows, visual feedback
- **Admin Control**: Content upload restricted to authorized users
- **Performance**: TanStack Query caching, atomic component reusability
- **Type Safety**: Comprehensive TypeScript with Zod validation