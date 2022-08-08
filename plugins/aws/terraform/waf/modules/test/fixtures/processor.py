from __future__ import print_function

import base64
import msgpack
import json

print('Loading function')


def lambda_handler(event, context):
  output = []

  for record in event['records']:
    payload = msgpack.unpackb(base64.b64decode(record['data']), raw=False)

    # Do custom processing on the payload here
    output_record = {
       'recordId': record['recordId'],
       'result': 'Ok',
       'data': base64.b64encode(json.dumps(payload).encode('utf-8') + b'\n').decode('utf-8')
    }
    output.append(output_record)

  print('Successfully processed {} records.'.format(len(event['records'])))
  return {'records': output}