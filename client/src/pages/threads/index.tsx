import Thread from "../../../components/Thread";

let threadsArray = [
    {id: 1, title: "kekw", text: "textBlablba", userId: 1, channelId: 2},
    {id: 2, title: "hahah", text: "blablal", userId: 1, channelId: 2},
    {id: 3, title: "abobus", text: "wowoow", userId: 2, channelId: 2},
]

const ThreadsPage = () => {
  const listThreads = threadsArray.map(thread => 
    <li key={thread.id}>
      <Thread {...thread}/>
    </li>
  );

  return (
    <>
      {listThreads}
    </>
  );
}

export default ThreadsPage;