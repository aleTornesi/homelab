-- name: GetVideos :many
SELECT *
FROM videos
LIMIT $1;

-- name: GetVideosFromId :many
SELECT *
FROM videos
WHERE id >= $2
LIMIT $1;

-- name: GetVidoById :one
SELECT *
FROM videos
WHERE id = $1;

-- name: GetPhotos :many
SELECT *
FROM photos
LIMIT $1;

-- name: GetPhotosFromId :many
SELECT *
FROM photos
WHERE id >= $2
LIMIT $1;

-- name: GetPhotoById :one
SELECT *
FROM photos
WHERE id = $1;

-- name: GetSongs :many
SELECT *
FROM songs
LIMIT $1;

-- name: GetSongsFromId :many
SELECT *
FROM songs
WHERE id >= $2
LIMIT $1;

-- name: GetSongById :one
SELECT *
FROM songs
WHERE id = $1;
