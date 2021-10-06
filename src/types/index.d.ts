interface Account {
  id: number;
  username: string;
  avatar_url: string;
  created_at: string;
}

interface Gallery {
  id: string;
  title: string;
  account: Account;
  media: Media[];
  tags: Tag[];
  cover: Media;
}

type MediaMimeType = 'image/jpeg' | 'image/png' | 'image/gif';
type MediaType = 'image';
type MediaExt = 'jpeg' | 'png' | 'gif';

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
  updated_at: "2021-10-01T00:08:51Z",
  deleted_at: null,
  next: null; //?
  comments: Comment[];
  account: {
    id: number;
    username: string;
    avatar: string;
  }
}

