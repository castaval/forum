'use client'

import Thread from "../../components/Thread"

type ThreadsData = {
  metadata: any;
  threads: any[];
};

const getThreadsData = async () => {
  const res = await fetch('http://localhost:4000/v1/threads')

  if (!res.ok) {
    throw new Error('Failed to fetch threads data');
  }

  const resjson = await res.json()
  console.log(resjson)

  return resjson;
}

const ThreadsPage = async () => {
  const threadsData: ThreadsData = await getThreadsData();

  const listThreads = threadsData.threads.map(thread => 
      <Thread {...thread}/>
  );

  return (
    <>
      {listThreads}
    </>
  );
}

export default ThreadsPage;