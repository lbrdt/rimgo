interface Account {
  id: number;
  username: string;
  avatar_url: string;
  created_at: string;
}

type MediaMimeType = 'image/jpeg' | 'image/png' | 'image/gif';
type MediaType = 'image';
type MediaExt = 'jpeg' | 'png' | 'gif';
type Sorting = 'newest' | 'oldest' | 'best' | 'viral';

interface Tag {
  tag: string;
  display: string;
  background_id: string;
  accent: string;
  is_promoted: boolean;
}

interface Media {
  id: string;
  account_id: number;
  mime_type: MediaMimeType;
  type: MediaType;
  name: string;
  basename: string;
  url: string;
  ext: MediaExt;
  width: number;
  height: number;
  size: number;
  metadata: {
    title: string;
    description: string;
    is_animated: boolean;
    is_looping: boolean;
    duration: number;
    has_sound: boolean;
  },
  created_at: string;
  updated_at: string | null;
}

type MediaPlatform = 'ios' | 'android' | 'api' | 'web';
interface Comment {
  id: number;
  parent_id: number;
  comment: string;
  account_id: number;
  post_id: string;
  upvote_count: number;
  downvote_count: number;
  point_count: number;
  vote: null; // ?
  platform_id: number;
  platform: MediaPlatform;
  created_at: string;
  updated_at: string;
  deleted_at: null;
  next: null; //?
  comments: Comment[];
  account: {
    id: number;
    username: string;
    avatar: string;
  }
}

interface Gallery {
  id: string;
  title: string;
  account: Account;
  media: Media[];
  tags: Tag[];
  cover: Media;
}

interface PostTag {
  name: string;
  display_name: string;
  followers: number;
  total_items: number;
  following: boolean;
  is_whitelisted: boolean;
  background_hash: string;
  thumbnail_hash: string;
  accent: string;
  background_is_animated: boolean;
  thumbnail_is_animated: boolean;
  is_promoted: boolean;
  description: string;
  logo_hash: string;
  logo_destination_url: string;
  description_annotations: {}
}

interface Post {
  id: string;
  account_id: number;
  type: MediaMimeType;
  animated: boolean;
  width: number;
  height: number;
  size: number;
  views: number;
  bandwidth: number;
  vote: null;
  favorite: boolean;
  nsfw: boolean;
  section: string;
  account_url: string;
  is_ad: boolean;
  in_most_viral: boolean;
  has_sound: boolean;
  tags: PostTag[];
  link: string;
  comment_count: number;
  ups: number;
  downs: number;
  score: number;
  points: number;
  is_album: boolean;
}

interface TagResult extends PostTag {
  items: Post[];
  success: boolean;
  status: number;
}

interface UserResult {
  id: number;
  username: string;
  bio: string;
  reputation_count: number;
  reputation_name: string;
  avatar_id: string;
  avatar_url: string;
  cover_id: string;
  cover_url: string;
  created_at: string;
}
