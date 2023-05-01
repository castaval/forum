type threadProps = {
    id: number;
    title: string;
    text: string;
    userId: number;
    channelId: number;
}

const Thread = (thread: threadProps) => {
    return (
        <div>
            {thread.title}
        </div>
    );
}

export default Thread;