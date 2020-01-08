#include "src/carnot/udf_exporter/udf_exporter.h"
#include "src/carnot/funcs/funcs.h"

namespace pl {
namespace carnot {
namespace udfexporter {

StatusOr<std::unique_ptr<compiler::RegistryInfo>> ExportUDFInfo() {
  auto registry = std::make_unique<udf::Registry>("udf_registry");

  funcs::RegisterFuncsOrDie(registry.get());

  udfspb::UDFInfo udf_proto = registry->ToProto();
  auto registry_info = std::make_unique<compiler::RegistryInfo>();
  PL_RETURN_IF_ERROR(registry_info->Init(udf_proto));
  return registry_info;
}

}  // namespace udfexporter
}  // namespace carnot
}  // namespace pl
