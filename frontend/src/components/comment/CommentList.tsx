import type { Comment as CommentType } from '../../types/comment';
import { formatDateTime } from '../../utils/format';

function CommentItem({ comment, depth = 0 }: { comment: CommentType; depth?: number }) {
  return (
    <div style={{ padding: '12px 0', borderBottom: '1px solid #f0f0f0', marginLeft: depth > 0 ? 24 : 0 }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 6 }}>
        <span style={{ fontWeight: 500, fontSize: 14 }}>{comment.author_name}</span>
        {comment.is_admin && (
          <span style={{ color: '#1677ff', fontSize: 11, background: '#e6f4ff', padding: '1px 6px', borderRadius: 4 }}>博主</span>
        )}
        <span style={{ color: '#999', fontSize: 12 }}>{formatDateTime(comment.created_at)}</span>
      </div>
      <p style={{ whiteSpace: 'pre-wrap', fontSize: 14, lineHeight: 1.6, color: '#333' }}>{comment.content}</p>
      {comment.children?.map((child) => (
        <CommentItem key={child.id} comment={child} depth={depth + 1} />
      ))}
    </div>
  );
}

interface Props {
  comments: CommentType[];
  onReply: (comment: CommentType) => void;
}

export default function CommentList({ comments }: Props) {
  return (
    <div>
      {comments.map((comment) => (
        <CommentItem key={comment.id} comment={comment} />
      ))}
    </div>
  );
}
